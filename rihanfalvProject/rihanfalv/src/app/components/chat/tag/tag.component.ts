import { Component, OnInit } from '@angular/core';

import { DataService } from '../../../services/data.service';
import { fromEvent } from 'rxjs'

@Component({
  selector: 'app-tag',
  templateUrl: './tag.component.html',
  styleUrls: ['./tag.component.css']
})
export class TagComponent implements OnInit {

  private inputValue: string = "";
  private resultData;

  private isMax: boolean = true;
  private nowData: any = [];

  private isBottom: boolean = false;
  private warn: string;

  constructor(
    private dataService: DataService,
  ) { }

  ngOnInit() {
    fromEvent(window, 'scroll').subscribe(() => {
      const h: any = document.documentElement.clientHeight;
      const H: any = document.body.clientHeight;
      const scrollTop: any = document.documentElement.scrollTop || document.body.scrollTop;
      if (h + scrollTop + 20 > H) {
        if (!this.isBottom) {
          setTimeout(() => { this.more()}, 1000);
        }
        this.isBottom = true;
      } else {
        this.isBottom = false;
      }
    });
  }

  tag(str: string) {
    this.inputValue += str + " ";
  }

  search() {
    // 关键字为空
    if (!this.inputValue) {
      this.isMax = true;
      this.warn = "请输入关键字后再进行搜索。";
      this.nowData = [];
      this.resultData = null;
      return
    }
    this.dataService.search(this.inputValue).subscribe((response) => {
      this.resultData = response;
      this.warn = null;
      if (this.resultData) {
        this.nowData = this.resultData.slice(0, 2);
        if (this.nowData.length == this.resultData.length) {
          this.isMax = true;
        } else {
          this.isMax = false;
        }
      } else { // 关键字找不到结果
        this.nowData = [];
        this.isMax = true;
      }
    });
  }

  more() {
    if (this.resultData) {
      this.nowData = this.resultData.slice(0, this.nowData.length + 5);
      if (this.nowData.length == this.resultData.length) {
        this.isMax = true;
      }
    }
  }
}
