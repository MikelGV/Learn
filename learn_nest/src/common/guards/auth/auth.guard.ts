import { CanActivate, ExecutionContext, Injectable } from '@nestjs/common';
import { rejects } from 'assert';
import { resolve } from 'path';
import { Observable } from 'rxjs';

@Injectable()
export class AuthGuard implements CanActivate {
  canActivate(
    context: ExecutionContext,
  ): boolean | Promise<boolean> | Observable<boolean> {
    const request =  context.switchToHttp().getRequest();
    return this.validateRequest(request);
  }

  async validateRequest(execContext: ExecutionContext): Promise<boolean> {
    const request = execContext.switchToHttp().getRequest();
    const user = request.user;
    if (user) {
      return true;
    }
  }

}
