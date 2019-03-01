import { Component, OnInit } from '@angular/core';

import { DataService } from '../../../services/data.service';

@Component({
  selector: 'app-tag',
  templateUrl: './tag.component.html',
  styleUrls: ['./tag.component.css']
})
export class TagComponent implements OnInit {

  private inputValue: string = "";
  private resultData;

  private isMax: boolean = true;
  private nowData: any;

  constructor(
    private dataService: DataService,
  ) { }

  ngOnInit() {
  }

  tag(str: string) {
    this.inputValue += str + " ";
  }

  search() {
    this.dataService.search(this.inputValue).subscribe((response) => {
      this.resultData = response;
      this.nowData = this.resultData.slice(0, 2);
      console.log(this.resultData);
      if (this.nowData.length == this.resultData.length) {
        this.isMax = true;
      } else {
        this.isMax = false;
      }
    });
  }

  more() {
    this.nowData = this.resultData.slice(0, this.nowData.length + 5);
    console.log(this.nowData);
    if (this.nowData.length == this.resultData.length) {
      this.isMax = true;
    }
  }
}
