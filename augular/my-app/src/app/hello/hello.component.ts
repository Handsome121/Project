import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-hello',
  templateUrl: './hello.component.html',
  styleUrls: ['./hello.component.css']
})
export class HelloComponent implements OnInit {
  name: string = 'Tina'
  box: string = 'div-dom'
  itemclass: string = 'item-p'
  h3Dom: boolean = true
  h3class: string = 'h3-dom font-w string'
  data: Date = new Date
  isActive: boolean = true
  constructor() { }

  ngOnInit() {
  }

}
