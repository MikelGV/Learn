import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { NestExpressApplication } from '@nestjs/platform-express';
import * as dotenv from "dotenv";
import { ValidationPipe } from './common/pipes/validation/validation.pipe';
import { RolesGuard } from './common/guards/roles/roles.guard';




dotenv.config();

async function bootstrap() {
  const app = await NestFactory.create<NestExpressApplication>(AppModule);
  app.useGlobalPipes(new ValidationPipe());
  app.useGlobalGuards(new RolesGuard());
  await app.listen(process.env.PORT);
}
bootstrap();
