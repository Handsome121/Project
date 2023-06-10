import { Component, OnInit } from "@angular/core";
import { FormControl, FormGroup } from "@angular/forms";
@Component({
  selector: "app-hello",
  templateUrl: "./hello.component.html",
  styleUrls: ["./hello.component.css"],
})
export class HelloComponent implements OnInit {
  name: string = "Tina";
  box: string = "div-dom";
  itemclass: string = "item-p";
  h3Dom: boolean = true;
  h3class: string = "h3-dom font-w string";
  data: Date = new Date();
  isActive: boolean = true;
  isMax: boolean = true;
  colors: Array<string> = ["red", "blue", "yellow", "green"];
  type: number = 3;
  title: string = "";

  age: FormControl = new FormControl("");

  loginForm: FormGroup = new FormGroup({
    userName: new FormControl(""),
    passWord: new FormControl(""),
  });

  constructor() {} //当前类的构造器

  ngOnInit() {}

  clickFun(e: Event) {
    console.log(e);
    alert("你点击了按钮");
  }
  inputChange(e: any) {
    console.log(e.target.value);
  }
  getUserName(v: string) {
    console.log(v);
  }
  ageChangeFun() {
    this.age.setValue(18);
  }
  submitFormFun() {
    console.log(this.loginForm.value);
  }
}
