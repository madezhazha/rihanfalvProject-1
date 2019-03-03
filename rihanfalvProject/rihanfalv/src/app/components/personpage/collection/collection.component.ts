import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-collection',
  templateUrl: './collection.component.html',
  styleUrls: ['./collection.component.css']
})
export class CollectionComponent implements OnInit {

  public flag:number=0;
  constructor() { }

  ngOnInit() {
  }

  thesis(){
    this.flag=0;
  }
  case(){
    this.flag=1;
  }
  topic(){
    this.flag=2;
  }

}
