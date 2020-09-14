package com.abigtomato.netty.simple;

import io.netty.buffer.ByteBuf;
import io.netty.buffer.Unpooled;
import io.netty.channel.Channel;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelInboundHandlerAdapter;
import io.netty.channel.ChannelPipeline;
import io.netty.util.CharsetUtil;

import java.util.concurrent.TimeUnit;

public class NettyServerHandler extends ChannelInboundHandlerAdapter {

    /**
     * 数据读取事件
     *
     * @param ctx 上下文对象，含有管道pipeline，通道channel
     * @param msg 客户端发送的数据
     */
    @Override
    public void channelRead(ChannelHandlerContext ctx, Object msg) {
        // 耗时业务解决方案1：用户程序自定义的普通任务 -> 异步执行 -> 提交到该channel对应的eventLoop的taskQueue中
        ctx.channel().eventLoop().execute(() -> {
            try {
                Thread.sleep(5 * 1000);
                ctx.writeAndFlush(Unpooled.copiedBuffer("hello world", CharsetUtil.UTF_8));
                System.out.printf("channel code: %s%n", ctx.channel().hashCode());
            } catch (Exception ex) {
                System.out.printf("发生异常: %s%n", ex.getMessage());
            }
        });

        // 耗时业务解决方案2：用户自定义定时任务 -> 异步执行 -> 该任务是提交到scheduleTaskQueue中
        ctx.channel().eventLoop().schedule(() -> {
            try {
                Thread.sleep(5 * 1000);
                ctx.writeAndFlush(Unpooled.copiedBuffer("hello world", CharsetUtil.UTF_8));
                System.out.printf("channel code: %s%n", ctx.channel().hashCode());
            } catch (Exception ex) {
                System.out.printf("发生异常: %s%n", ex.getMessage());
            }
        }, 5, TimeUnit.SECONDS);

        System.out.printf("服务器读取线程: %s, channel: %s%n", Thread.currentThread().getName(), ctx.channel());
        System.out.printf("服务器上下文: %s%n", ctx);

        Channel channel = ctx.channel();
        System.out.printf("客户端地址: %s%n", channel.remoteAddress());
        // pipeline本质是一个双向链接，有出站和入站操作
        ChannelPipeline pipeline = ctx.pipeline();

        // 将msg转成一个ByteBuf，ByteBuf是Netty提供的，不是NIO的ByteBuffer
        ByteBuf buf = (ByteBuf) msg;
        System.out.printf("客户端发送的消息: %s%n", buf.toString(CharsetUtil.UTF_8));
    }

    /**
     * 数据读取完毕事件
     *
     * @param ctx 上下文对象
     */
    @Override
    public void channelReadComplete(ChannelHandlerContext ctx) {
        // 将数据写入到缓存，并刷新
        ctx.writeAndFlush(Unpooled.copiedBuffer("hello world", CharsetUtil.UTF_8));
    }

    /**
     * 异常事件
     *
     * @param ctx 上下文
     * @param cause 异常
     */
    @Override
    public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) {
        ctx.close();
    }
}
