import { CanActivate, ExecutionContext, Injectable } from '@nestjs/common';
import { Reflector } from '@nestjs/core';
import { Roles } from 'src/common/decorators/roles/roles.decorator';
import { User } from 'src/user/entities/user.entity';

@Injectable()
export class RolesGuard implements CanActivate {
  constructor(private reflector: Reflector) {}

  async matchRoles(roles: string[], userRole: User['roles']) {
    let match = false;

    if (roles.indexOf(userRole) > -1) {
      match = true;
    }

    return match;
  }

  async canActivate(context: ExecutionContext) {
    const roles = this.reflector.get(Roles, context.getHandler());
    if (!roles) {
      return true;
    }
    const request = context.switchToHttp().getRequest();
    const user = request.user;
    return this.matchRoles(roles, user.roles);
  }
}
