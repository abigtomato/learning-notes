import hashlib
import json
from time import time
from typing import Any, Dict, List, Optional
from urllib.parse import urlparse
from uuid import uuid4

import requests
from flask import Flask, jsonify, request


class Blockchain:
    def __init__(self):
        # 存储当前未被打包成区块的交易
        self.current_transactions = []
        # 存储链上的所有区块
        self.chain = []
        # 存储所有的节点信息
        self.nodes = set()
        # 初始化创世块
        self.new_block(previous_hash='1', proof=100)

    def register_node(self, address: str) -> None:
        """
        Add a new node to the list of nodes 注册节点
        :param address: http://192.168.121.100:5000
        :return: None
        """
        parse_url = urlparse(address) # 解析url
        self.nodes.add(parse_url.netloc) # 存入node列表

    def new_block(self, previous_hash: Optional[str], proof: int) -> Dict[str, Any]:
        """
        打包current交易生成新区块
        :param previous_hash: 上一个区块的Hash
        :param proof: 工作量
        :return: block格式字典
        """
        block = {
            'index': len(self.chain) + 1, # 当前区块高度
            'timestamp': time(), # 时间戳
            'transactions': self.current_transactions, # 打包的交易
            'proof': proof, # 工作量
            'previous_hash': previous_hash or self.hash(self.chain[-1]), # 参数传入或计算链上最后一个区块的Hash
        }

        self.current_transactions = [] # 清空打包后的交易
        self.chain.append(block) # 添加新区块到链上

        return block

    def new_transaction(self, sender: str, recipient: str, amount: int) -> int:
        """
        生成一笔新的交易，等待下一次打包生成新区块
        :param sender: 发起方
        :param recipient: 接收方
        :param amount: BTC金额
        :return:
        """
        transaction = {
            'sender': sender,
            'recipient': recipient,
            'amount': amount,
        }

        self.current_transactions.append(transaction)

        return self.last_block['index'] + 1

    @property
    def last_block(self) -> Dict[str, Any]:
        """
        返回最后一个区块
        :return: 区块结构
        """
        return self.chain[-1]

    @staticmethod
    def hash(block: Dict[str, Any]) -> str:
        """
        计算区块的Hash值
        :param block:
        :return:
        """
        block_string = json.dumps(block, sort_keys=True).encode()
        return hashlib.sha256(block_string).hexdigest()

    def proof_of_work(self, last_proof: int):
        """
        共识机制: 工作量证明
        :param last_proof: 链上最后一个区块的工作量
        :return: 挖矿的工作量
        """
        proof = 0 # 随机数做为本次工作量的起始
        # 对last_proof和不断变化的随机数拼接进行Hash运算，直到满足系统难度值，挖矿才会成功
        while self.valid_proof(last_proof, proof) is False:
            proof += 1 # 不断变化的随机数
        return proof

    @staticmethod
    def valid_proof(last_proof: int, proof: int):
        """
        验证本次挖矿是否符合系统难度值
        :param last_proof: 链上最后一个区块的工作量
        :param proof: 本次挖矿随机数
        :return: 是否成功
        """
        guess = f'{last_proof}{proof}'.encode()
        guess_hash = hashlib.sha256(guess).hexdigest()
        return guess_hash[:4] == '0000' # 难度值由系统决定，前置0越多，越复杂

    def resolve_conflicts(self) -> bool:
        """
        解决区块链网络中的分叉问题(使用网络中最长的链)
        :return: 如果链被取代返回True，反之False
        """
        neighbours = self.nodes
        new_chain = None
        max_length = len(self.chain)

        for node in neighbours:
            # 通过web服务获取其他节点保存的区块链结构
            response = requests.get(f'http://{node}/chain')

            if response.status_code == 200:
                length = response.json()['length']
                chain = response.json()['chain']

                # 判断区块链的合法性和链的长度
                if length > max_length and self.valid_chain(chain):
                    max_length = length
                    new_chain = chain

        # 若链长不同，则取网络中最长的链同步
        if new_chain:
            self.chain = new_chain
            return True

        return False

    def valid_chain(self, chain: List[Dict[str, Any]]) -> bool:
        """
        验证区块链的合法性
        :param chain:
        :return:
        """
        prev_block = chain[0]
        current_index = 1

        while current_index < len(chain):
            block = chain[current_index]

            # 验证条件: 后一个区块的previous_hash等于前一个区块的Hash
            if block['previous_hash'] != self.hash(prev_block):
                return False

            # 验证条件: Hash(前一个区块的proof+当前区块的proof)符合系统难度值
            if not self.valid_proof(prev_block['proof'], block['proof']):
                return False

            prev_block = block
            current_index += 1

        return True


app = Flask(__name__)
node_identifier = str(uuid4()).replace('-', '')
bc = Blockchain()


@app.route('/transactions/new', methods=['POST'])
def new_transaction():
    """
    新交易请求的controller
    :return: response body和response code
    """
    values = request.get_json() # 获取表单数据(JSON映射成Dict)

    # 判断表单数据是否符合要求
    if not all(e in values for e in ['sender', 'recipient', 'amount']):
        return 'Missing values', 400

    # 生成新的交易
    index = bc.new_transaction(values['sender'], values['recipient'], values['amount'])

    return jsonify({
        'message': f'Transaction will be added to Block {index}'
    }), 201


@app.route('/mine', methods=['GET'])
def mine():
    """
    发起挖矿生成新区块
    :return: response body和response code
    """
    last_proof = bc.last_block['proof']
    proof = bc.proof_of_work(last_proof) # 挖矿
    block = bc.new_block(None, proof)    # 打包新的区块

    # 生成一笔挖矿交易
    bc.new_transaction(
        sender='0',
        recipient=node_identifier,
        amount=12,
    )

    return jsonify({
        'message': 'New Block Forged',
        'index': block['index'],
        'transactions': block['transactions'],
        'proof': block['proof'],
        'previous_hash': block['previous_hash'],
    }), 200


@app.route('/chain', methods=['GET'])
def full_chain():
    """
    获取整条区块链
    :return: response body和response code
    """
    return jsonify({
        'chain': bc.chain,
        'length': len(bc.chain),
    }), 200


@app.route('/nodes/register', methods=['POST'])
def register_nodes():
    """
    注册新的节点
    :return: response body和response code
    """
    values = request.get_json()

    nodes = values.get('nodes')
    if nodes in None:
        return 'Error: Please supply a valid list of nodes', 400

    for node in nodes:
        bc.register_node(node)

    return jsonify({
        'message': 'New nodes have been added',
        'total_nodes': list(bc.nodes),
    }), 201


@app.route('/nodes/resolve', methods=['GET'])
def consensus():
    """
    解决区块链网络中链的分叉问题
    :return:
    """
    replaced = bc.resolve_conflicts()

    if replaced:
        response = {
            'message': 'Our chain was replaced',
            'new_chain': bc.chain,
        }
    else:
        response = {
            'message': 'Out chain is authoritative',
            'chain': bc.chain,
        }

    return jsonify(response), 200


if __name__ == '__main__':
    from argparse import ArgumentParser

    # 获取命令行参数
    parser = ArgumentParser()
    parser.add_argument('-p', '--port', default=5000, type=int, help='port to listen on')
    args = parser.parse_args()

    app.run(host='127.0.0.1', port=args.port)