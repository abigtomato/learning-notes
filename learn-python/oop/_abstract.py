import abc


class CacheBase(metaclass=abc.ABCMeta):
    @abc.abstractmethod
    def get(self):
        pass

    @abc.abstractmethod
    def set(self, key, value):
        pass


class RedisCache(CacheBase):
    def get(self):
        pass

    def set(self, key, value):
        pass


if __name__ == '__main__':
    redis = RedisCache()
    print(redis)

    print(isinstance(redis, CacheBase))