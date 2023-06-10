import { Component } from "@angular/core";

@Component({
  selector: "app-root", //选择器
  templateUrl: "./app.component.html", //对应组件模板
  styleUrls: ["./app.component.css"], //需要加载到当前模块的样式表
})
export class AppComponent {
  title = "app works!";
}
