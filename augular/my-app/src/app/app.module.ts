import { BrowserModule } from "@angular/platform-browser";
import { NgModule } from "@angular/core";
import { FormsModule } from "@angular/forms";
import { HttpModule } from "@angular/http";
import { ReactiveFormsModule } from "@angular/forms";

import { AppComponent } from "./app.component";
import { HelloComponent } from "./hello/hello.component";

@NgModule({
  //声明组件内用到的视图
  declarations: [AppComponent, HelloComponent],
  //声明组件需要用到的类
  imports: [BrowserModule, FormsModule, HttpModule, ReactiveFormsModule],
  //全局服务
  providers: [],
  //根组件
  bootstrap: [AppComponent],
})
export class AppModule {}
