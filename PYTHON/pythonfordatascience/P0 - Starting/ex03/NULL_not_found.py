def NULL_not_found(object: any) -> int:
    match object:
        case None:
            print("Nothing: None", type(object))
        case float() if object != object:  # Vérification pour NaN
            print("Cheese: nan", type(object))
        case bool():
            print("Fake: False", type(object))
        case int():
            print("Zero: 0 ", type(object))
        case str() if object == "":
            print("Empty:", type(object))
        case _:
            print("Type not Found")
    return 1
