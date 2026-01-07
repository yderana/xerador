import { Column, Entity, PrimaryGeneratedColumn } from "typeorm";

@Entity()
export class Template {
    @PrimaryGeneratedColumn()
    id: number;

    @Column()
    name: string;

    @Column({ name: "created_by" })
    createdBy: string;

    @Column({ name: "created_at" })
    createdAt: Date;

    @Column({ name: "modified_by", nullable: true })
    modifiedBy?: string;

    @Column({ name: "modified_at", nullable: true })
    modifiedAt?: Date;
}
