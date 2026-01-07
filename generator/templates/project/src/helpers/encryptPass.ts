import * as bcrypt from 'bcrypt';

const saltOrRounds = 10;
export async function encryptPassword(password: string) {
    return await bcrypt.hash(password, saltOrRounds);
}

export async function validatePassword(password: string, hash: string) {
    return await bcrypt.compare(password, hash);
}