class node: #plant
    def __init__(self,val=None):
        self.thirst = 0
        self.plant = {val : self.thirst}
        self.nx = None

class linkedlist: #fields
    def __init__(self):
        self.type = {"dirt" : "#","water" : "~"}
        self.fld = []
        self.frst = None
        self.nx = None

    def new_fld(self,frm):
        dirt = node(self.type["dirt"])
        amt = len(frm)
        print("making new field number:",amt + 1)
        frm.append([dirt])
        
    def add_plnt(self,frm,rw=None):
        if rw is None:
            rw = 0
        
        print("add plant to row:",rw + 1)
        fake = node(self.type["dirt"])

        self.frst = frm[rw][0]
        start = self.frst

        while start.nx:
            start = start.nx

        start.nx = fake #give the last node and new node
        frm[rw].append(start.nx)

    def water_plnt(self,frm,rw,num):
        self.frst = frm[rw][0]
        start = self.frst
        cnt = 1
        
        while start.nx:
            if cnt != num:
                cnt += 1
                start = start.nx
                print(cnt,":",num)

            if self.type["water"] in frm[rw][num-1].plant.keys():
                print("no point checking mud for more water")
            else:
                frm[rw][num-1].plant[self.type["water"]] = frm[rw][num-1].plant[self.type["dirt"]]
                del frm[rw][num-1].plant[self.type["dirt"]]
                frm[rw][num-1].plant[self.type["water"]] = 100

            
            if frm[rw][num-1].thirst <= 50:
                print("checking the soil lvl if it needs water:")
                tmp = list(start.plant) #outputs a list of keys
                start.plant[tmp[0]] = 100
                break



    def plant(self,frm,rw,num,seed):
        self.frst = frm[rw][0]
        start = self.frst
        cnt = 1
        myseed = farm()
        
        while start.nx:
            if cnt != num:
                cnt += 1
                start = start.nx
                print(cnt,":",num)
            
            if self.type["water"] not in frm[rw][num-1].plant.keys():
                if self.type["dirt"] in frm[rw][num-1].plant.keys():
                    print("running need water first")

            if frm[rw][num-1].plant[self.type["water"]]:
                print("checking the soil lvl if it needs water:")
                start.plant[myseed.seeds[seed]] = start.plant[self.type["water"]]
                del start.plant[self.type["water"]]
                break

class farm:
    def __init__(self):
        self.plot = {}
        self.inventory = {}
        self.seeds = {"sunflower" : "+","corn" : "c"}

    def display(self,frm):
        for i,x in enumerate(frm):
            print(i,[(k.plant) for k in x])


class clock():
    def __init__(self):
        self.day = 30
        self.tick = 1

    def actions():

    def simulate():

fld1 = linkedlist()
myfrm = farm()
corn = node("c")

frm = []
print("orignal farm",frm)

"""
[*] add new row to farm 
[*] add new plant to given row with nodes
"""
fld1.new_fld(frm)
print("farm with its first field",frm)
fld1.add_plnt(frm,0)
print("putting plant in the field {node}",[(x.plant) for x in frm[0]])
fld1.new_fld(frm)
print("secound row of farm with first plant",frm)
fld1.add_plnt(frm,1)
print("putting plant in the field {node}",[(x.plant) for x in frm[1]])
print("####################")
myfrm.display(frm)

"""
[*] water plants check if mud
[*] plant certain seed
"""
fld1.water_plnt(frm,0,2)
print("water plant: ",[ (x.plant) for x in frm[0] ])
fld1.water_plnt(frm,0,2)
print("water the same plant and make sure you cant instead just check its thirst")
print("water plant: ",[ (x.plant) for x in frm[0] ])
print(myfrm.seeds.items())
fld1.plant(frm,0,2,"corn")
print("water the dirt spot: ",[ (x.plant) for x in frm[0]  ])
fld1.water_plnt(frm,0,1)
print("seems to have enough water: ",[ (x.plant) for x in frm[0]  ])
fld1.plant(frm,0,1,"sunflower")
print("lets put a seed: ",[ (x.plant) for x in frm[0]  ])
