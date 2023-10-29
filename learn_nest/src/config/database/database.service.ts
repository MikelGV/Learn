import { Injectable } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { TypeOrmModuleOptions, TypeOrmOptionsFactory } from '@nestjs/typeorm';

@Injectable()
export class DatabaseService implements TypeOrmOptionsFactory {

    constructor(private configService: ConfigService) {}

    createTypeOrmOptions(): TypeOrmModuleOptions {
        return {
            type: 'postgres',
            host: process.env.DB_HOST,
            port: parseInt(process.env.DB_PORT),
            password: process.env.DB_PWORD,
            username: process.env.DB_USERNAME,
            database: process.env.DB_DB,
            entities: [
            "dist/**/*.entity{.ts,.js}",
            ],
            synchronize: true,
        };
    }
}
