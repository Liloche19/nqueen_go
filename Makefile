SRC	=	src/main.go

NAME	=	nqueen

all:
	go build -o $(NAME) $(SRC)

clean:

fclean:	clean
	rm -f $(NAME)

re:	fclean all
