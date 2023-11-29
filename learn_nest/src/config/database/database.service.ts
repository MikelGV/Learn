import { Injectable, OnModuleInit, Scope } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { TypeOrmModuleOptions, TypeOrmOptionsFactory } from '@nestjs/typeorm';

@Injectable({ scope: Scope.DEFAULT })
export class DatabaseService implements TypeOrmOptionsFactory, OnModuleInit {
    onModuleInit() {
        console.log('this starts');
    }

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
