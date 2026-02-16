#[derive(Debug)]
#[warn(dead_code)]
struct Boss {
    life: i32,
    mana: i32,
    dead: bool,
    name: String,
    age: i32,
}

enum Shot {
    Critical(i32),
    Hit(i32), 
}

trait Attack { // Creating methods
    fn attack_head(self: &mut Self) -> Shot;
    fn attack_body(self: &mut Self) -> Shot;
    fn attack_leg(self: &mut Self) -> Shot;
}

impl Attack for Boss { // Implementing methods
    fn attack_head(self: &mut Self) -> Shot {
        let damage = 40;
        self.life -= damage;
        if self.life <= 0 {
            self.dead = true;
            self.life = 0;
            Shot::Critical(damage)
        } else {
            Shot::Hit(damage)
        }
    }

    fn attack_body(self: &mut Self) -> Shot {
        let damage = 20;
        self.life -= damage;
        if self.life <= 0 {
            self.life = 0;
            Shot::Critical(damage)
        } else {
            Shot::Hit(damage)
        }
    }

    fn attack_leg(self: &mut Self) -> Shot {
        let damage = 10;
        self.life -= damage;
        if self.life <= 0 {
            self.dead = true;
            self.life = 0;
            Shot::Critical(damage)
        } else {
            Shot::Hit(damage)
        }
    }
}

pub fn everything_about_structs_traits() {
    println!("\n....Structs, Traits, Impl, Enum....\n");
    let mut new_boss = Boss {
        life: 100,
        mana: 100,
        dead: false,
        name: String::from("The Reaper"),
        age: 28,
    };

    new_boss.attack_head();
    new_boss.attack_body();
    new_boss.attack_leg();
    new_boss.attack_head();

    println!("Here is boss after the attack {:?} ", new_boss);

    
}