import { Component, OnInit, Pipe, PipeTransform } from '@angular/core';

import { DataService } from '../../../services/data.service';
import { Router, NavigationExtras } from '@angular/router';

import { DomSanitizer } from '@angular/platform-browser';


@Component({
  selector: 'app-popular',
  templateUrl: './popular.component.html',
  styleUrls: ['./popular.component.css']
})
export class PopularComponent implements OnInit {

  private threadList: any;
  // 当前数据的数组是否是最大值
  private isMax: boolean = false;
  private nowData: Array<any> = [];


  // 韩国日本的标志
  private flag: any;

  constructor(
    private dataService: DataService,
    private router: Router,
    private sanitizer: DomSanitizer,
  ) { }

  ngOnInit() {
    this.dataService.getThreadList().subscribe((response) => {
      if (localStorage.getItem("JapanOrKorea") == "韩") {
        this.flag = 0;
      } else {
        this.flag = 1;
      }

      this.threadList = response;
      // console.log(this.threadList);
      if (this.threadList) {
        for (let i = 0; i < this.threadList.length; i++) {
          const element = this.threadList[i];
          if (element.user.Image.indexOf("assets") == -1) {
            let temp = 'data:image/png;base64, ' + element.user.Image; //给base64添加头缀
            element.user.Image = this.sanitizer.bypassSecurityTrustUrl(temp);
          }
          // console.log(element.user.Image);
        }
      }
      this.cutArray(this.getCountryData(), 2);
    });
  }

  read(item) {
    this.dataService.addVisitNum(item.thread.ID).subscribe(() => {
      let navigationExtras: NavigationExtras = {
        queryParams: {
          topicID: item.thread.ID,
        },
      }
      this.router.navigate(['/post'], navigationExtras);
    });
  }

  more() {
    this.cutArray(this.getCountryData(), 2);
  }

  getJapanKorea(isJapan: boolean) {
    if (isJapan) {
      this.flag = 1;
    } else {
      this.flag = 0;
    }

    this.nowData = [];
    this.cutArray(this.getCountryData(), 2);
  }

  getCountryData(): Array<any> {
    if (this.threadList){
      let countryData: Array<any> = [];
      for (let i = 0; i < this.threadList.length; i++) {
        const element = this.threadList[i];
        if (element.thread.Japanorkorea == this.flag) {
          countryData.push(element);
        }
      }
      return countryData;
    }
    return [];
  }

  // 截取展示数据，并且判断现在长度是否为最大
  cutArray(countryData: Array<any>, length: number) {
    if (countryData) {
      this.nowData = countryData.slice(0, this.nowData.length + length);
      if (this.nowData.length == countryData.length) {
        this.isMax = true;
      } else {
        this.isMax = false;
      }
    } else {
      this.isMax = true;
    }
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