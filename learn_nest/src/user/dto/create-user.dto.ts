import { IsEmail, IsNotEmpty, IsString, Matches, MinLength } from "class-validator";


const pwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*d)(?=.*[@$!%*?&])[A-Za-zd@$!%*?&]{8,20}$/;

export class CreateUserDto {
    @IsString()
    @MinLength(3, {message: 'Username must have at least 3 characters'})
    @IsNotEmpty()
    username: string;

    @IsEmail()
    @IsNotEmpty()
    email: string;
  
    @IsNotEmpty()
    @Matches(pwordRegex, {
        message: 
        `Password must contain a Minimum 8 and Maximum 20 characters,
        at least one uppercase letter,
        one lowercase letter,
        one number and
        one special character`,
    })
    password: string;
  }
