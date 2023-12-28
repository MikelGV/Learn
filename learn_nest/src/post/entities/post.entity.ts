import { Column, Entity, PrimaryGeneratedColumn } from "typeorm";

export class Post {
    @PrimaryGeneratedColumn()
    id: number

    @Column({ type:"varchar", length: 50 })
    title: string

    @Column({ type:"varchar", length: 5000 })
    description: string
}
