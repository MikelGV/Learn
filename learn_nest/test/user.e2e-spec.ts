import { INestApplication } from '@nestjs/common';
import { Test } from '@nestjs/testing';
import { UserModule } from '../src/user/user.module';
import { UserService } from '../src/user/user.service';
import * as request from 'supertest';

describe('User', () => {
  let app: INestApplication;
  let userService = {
    findAll: () => ['test'],
  };

  beforeAll(async () => {
    const moduleRef = await Test.createTestingModule({
      imports: [UserModule],
    })
      .overrideProvider(UserService)
      .useValue(userService)
      .compile();

    app = moduleRef.createNestApplication();
    await app.init();
  });

  it('/GET user', () => {
    return request(app.getHttpServer())
      .get('/user')
      .expect(200)
      .expect({ data: userService.findAll() });
  });
  afterAll(async () => {
    await app.close();
  });
});
