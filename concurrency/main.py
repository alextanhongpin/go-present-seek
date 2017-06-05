# import time
# from concurrent.futures import ThreadPoolExecutor


# def delay1000():
# 	time.sleep(1)
# 	return True

# def delay2000():
# 	time.sleep(2)
# 	return True


# executor = ThreadPoolExecutor(max_workers=2)
# start = time.time()
# a = executor.submit(delay1000())
# b = executor.submit(delay2000())
# end = time.time()
# print end - start