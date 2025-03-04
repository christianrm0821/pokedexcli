import turtle
import time

'''
turtle.forward()
turtle.backward()
turtle.right()
turtle.left()
turtle.penup()
turtle.pendown()
turtle.color()
turtle.shape()
turtle.stamp()
turtle.dot()
turtle.clear()
turtle.reset()
'''


def create_polygon(sides):
    degree = (sides-2)*180
    individual_degree = degree/sides
    for i in range(sides):
        turtle.forward(100)
        turtle.right(180.0-individual_degree)

turtle.exitonclick()

create_polygon(6)





