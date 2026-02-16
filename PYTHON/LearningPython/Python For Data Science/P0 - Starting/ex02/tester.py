from find_ft_type import all_thing_is_obj

ft_list = ["Hello", "toto"]
ft_tuple = ("Hello", "tutu")
ft_set =  {"Hello", "tata"}
ft_dict = {"Hello" : "titi"}

all_thing_is_obj(ft_list)
all_thing_is_obj(ft_tuple)
all_thing_is_obj(ft_set)
all_thing_is_obj(ft_dict)
all_thing_is_obj("Brian")
all_thing_is_obj("Toto")
print(all_thing_is_obj(10))