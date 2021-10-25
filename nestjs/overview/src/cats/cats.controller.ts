import {
  Controller,
  Get,
  Header,
  HttpCode,
  Param,
  Post,
  Req,
} from '@nestjs/common';
import { Request } from 'express';

@Controller('cats')
export class CatsController {
  @Get()
  findAll(@Req() request: Request): string {
    return 'hoge';
  }

  @Post()
  @HttpCode(204)
  @Header('Cache-Control', 'none')
  create(): string {
    return 'cat created.';
  }

  @Get(':id')
  findOne(@Param() params): string {
    console.log(params.id);
    return `This action returns a #${params.id} cat`;
  }
}
