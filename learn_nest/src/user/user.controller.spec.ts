import { Test, TestingModule } from '@nestjs/testing';
import { UserController } from './user.controller';
import { UserService } from './user.service';
// import { ModuleMocker, MockFunctionMetadata } from 'jest-mock';

// const moduleMocker = new ModuleMocker(global);

describe('UserController', () => {
  let userController: UserController;
  let userService: UserService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [UserController],
      providers: [
        {
          provide: UserService,
          useValue: {
            findAll: jest.fn().mockResolvedValue([
              {
                username: 'mikel',
                email: 'mikelgaldosv@gmail.com',
                roles: 'admin',
                password: '78948393m',
                id: 1,
              },
            ]),
          },
        },
      ],
    })
      // .useMocker((token) => {
      //   const results = ['test1', 'test2'];
      //   if (token === UserService) {
      //     return { findAllUsers: jest.fn().mockResolvedValue(results) };
      //   }
      //   if (typeof token === 'function') {
      //     const mockMetadata = moduleMocker.getMetadata(
      //       token,
      //     ) as MockFunctionMetadata<any, any>;
      //     const Mock = moduleMocker.generateFromMetadata(mockMetadata);
      //     return new Mock();
      //   }
      // })
      .compile();
    userController = module.get<UserController>(UserController);
    userService = module.get<UserService>(UserService);
  });

  describe('findAll', () => {
    it('should return an array of users', async () => {
      const result = [
        {
          username: 'mikel',
          email: 'mikelgaldosv@gmail.com',
          roles: 'admin',
          password: '78948393m',
          id: 1,
        },
      ];
      jest.spyOn(userService, 'findAll').mockResolvedValue(result);
      expect(await userService.findAll()).toBe(result);
    });
  });
});
