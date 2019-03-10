import { Component, OnInit, Pipe, PipeTransform } from '@angular/core';

import { DataService } from '../../../services/data.service';
import { Router, NavigationExtras } from '@angular/router';


@Component({
  selector: 'app-popular',
  templateUrl: './popular.component.html',
  styleUrls: ['./popular.component.css']
})
export class PopularComponent implements OnInit {

  private threadList: any;
  private isMax: boolean = false;
  private nowData: Array<any> = [];

  // 韩国日本的标志
  private flag: any;

  constructor(
    private dataService: DataService,
    private router: Router
  ) { }

  ngOnInit() {
    this.dataService.getThreadList().subscribe((response) => {
      if (localStorage.getItem("JapanOrKorea") == "韩") {
        this.flag = 0;
      } else {
        this.flag = 1;
      }
      // debugger;
      this.threadList = response;
      if (this.threadList) {
        this.nowData = this.threadList.slice(0, 1);
        if (this.nowData.length == this.threadList.length) {
          this.isMax = true;
        } else {
          this.isMax = false;
        }
      }
      // console.log(this.threadList);
      // console.log(this.nowData);
    });
  }

  read(item) {
    let navigationExtras: NavigationExtras = {
      queryParams: {
        topicID: item.thread.ID,
      },
    }
    this.router.navigate(['/post'], navigationExtras)
  }

  more() {
    if (this.threadList != null) {
      this.nowData = this.threadList.slice(0, this.nowData.length + 5);
      if (this.nowData.length == this.threadList.length) {
        this.isMax = true;
      }
    } else {
      this.isMax = true;
    }
  }

  getJapanKorea(isJapan: boolean) {
    if (isJapan) {
      this.flag = 1;
    } else {
      this.flag = 0;
    }
  }

  getLoginData($event){
    
  }

}

@Pipe({
  name: 'myPipe',
})
export class myPipe implements PipeTransform {
  transform(data: Array<any>, args?: string): Array<any> {
    let newArray: Array<any> = [];
    for (let i = 0; i < data.length; i++) {
      const element = data[i];
      if (element.thread.Japanorkorea === args) {
        newArray.push(element);
      }
    }
    return newArray;
  }
}