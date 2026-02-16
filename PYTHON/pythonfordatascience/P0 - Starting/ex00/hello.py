def hello_world():
    ft_list = ["Hello", "toto"]
    ft_tuple = ("Hello", "tutu")
    ft_set =  {"Hello", "tata"}
    ft_dict = {"Hello" : "titi"}

    ft_list[1] = "World"
    print(ft_list)

    x = list(ft_tuple)
    x[1] = "France"
    ft_tuple = tuple(x)
    print(ft_tuple)

    print("A set is a collection which is unordered, unchangeable*, and unindexed. So the result can be weird sometimes")
    print(ft_set)
    ft_set.discard("tata")
    Paris = {"Paris"}
    ft_set.update(Paris)
    print(ft_set)
    
    ft_dict["Hello"] = "42Paris"
    print(ft_dict)


hello_world()

