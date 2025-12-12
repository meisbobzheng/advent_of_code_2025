import math

def solve():
    try:
        file_path = "./values.txt"
        #file_path = "./sample.txt"
        with (open(file_path, 'r') as file) :
            
            count = 0
            current = 50

            for line in file:
                first = line[0]
                rotation = int(line[1:])
                clicks_to_first = 0
                print(f"Current position: {current} | Input: {line}")

                if first == "R":
                    clicks_to_first = (100 - current) % 100 
                    current = current + rotation

                    
                else:
                    clicks_to_first = current
                    current = current - rotation

                if clicks_to_first == 0:
                    clicks_to_first = 100

                current = current % 100
                print(f"New position: {current}")

                zeros = 0

                if rotation >= clicks_to_first:
                    print("GETS FIRST CLICK")
                    print(f"rotation: {rotation} | clicks to first: {clicks_to_first}")
                    rotation = rotation - clicks_to_first
                    zeros = zeros + 1


                click_floor = max(math.floor(rotation / 100), 0)
                zeros = zeros + click_floor
                count += zeros
                print(f"Additional clicks: {click_floor}")
                print(f"Total clicks: {zeros}")
                print(f"Trailing total: {count}")
                print("=====")
                

            print(count)


    except Exception as e:
        print(f"Error: {e}")
    

if __name__ == "__main__":
    solve()
