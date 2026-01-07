import { Module, HttpModule } from "@nestjs/common";
import { TypeOrmModule } from "@nestjs/typeorm";
import { TemplateService } from "@template/providers/template.service";
import { Template } from "@template/repository/template.entity";
import { TemplateController } from "@template/controllers/template.controller";
import { TemplateResolvers } from "@template/controllers/template.resolvers";
import { TemplateRepository } from "@template/repository/template.repository";

@Module({
  imports: [TypeOrmModule.forFeature([Template]), HttpModule],
  providers: [TemplateService, TemplateResolvers, TemplateRepository],
  controllers: [TemplateController]
})
export class TemplateModule { }
