x = foo
y = $(x) bar
x = later

all:
	echo $(y)
