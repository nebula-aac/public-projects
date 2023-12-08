import { Container } from "inversify";
import { TYPES } from "./types";
import { Warrior, Weapon, ThrowableWeapon } from "./interfaces";
import { Ninja, Katana, Shuriken } from "./entities";

const newContainer = new Container();
newContainer.bind<Warrior>(TYPES.Warrior).to(Ninja);
newContainer.bind<Weapon>(TYPES.Weapon).to(Katana);
newContainer.bind<ThrowableWeapon>(TYPES.ThrowableWeapon).to(Shuriken);

export { newContainer }