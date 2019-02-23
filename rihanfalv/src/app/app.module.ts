import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { ReactiveFormsModule } from '@angular/forms';
import { FormsModule } from '@angular/forms';
import { HttpClientModule} from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { FeedbackComponent } from './components/feedback/feedback.component';
import { PaperComponent } from './components/paper/paper.component';
import { PaperwebComponent } from './components/paper/paperweb/paperweb.component';
import { SearchComponent } from './components/search/search.component';
import { CaseComponent } from './components/case/case.component';
import { PersonpageComponent } from './components/personpage/personpage.component';
import { ChatComponent } from './components/chat/chat.component';
import { HeadComponent } from './components/head/head.component';
import { HomepageComponent} from './components/homepage/homepage.component';

import { FeedbacksuccessComponent } from './components/Feedback/feedbacksuccess/feedbacksuccess.component';
import { PersonalfeedbackComponent } from './components/Feedback/personalfeedback/personalfeedback.component';

// 头部组件
import { LangingComponent } from './components/head/langing/langing.component';
import { LandComponent } from './components/head/langing/land/land.component';
import { RegisterComponent } from './components/head/langing/register/register.component';
import { ForgetpasswordComponent } from './components/head/langing/forgetpassword/forgetpassword.component';
import { WebheadComponent } from './components/head/webhead/webhead.component';
import { UploadHeadimageComponent } from './components/personpage/upload-headimage/upload-headimage.component';
// 个人主页
import { UserComponent } from './components/personpage/user/user.component' ;
import { CollectionComponent } from './components/personpage/collection/collection.component' ;
import { ThesisComponent } from './components/personpage/collection/thesis/thesis.component' ;
import { DatePipe } from '@angular/common';
// 服务
import {GetdataService} from '../app/services/getdata.service';
//

@NgModule({
  declarations: [
    AppComponent,
    FeedbackComponent,
    PaperComponent,
    PaperwebComponent,
    SearchComponent,
    CaseComponent,
    PersonpageComponent,
    ChatComponent,
    HeadComponent,
    HomepageComponent,
    // 头部组件
    LangingComponent,
    LandComponent,
    RegisterComponent,
    ForgetpasswordComponent,
    WebheadComponent,
    UserComponent,
    CollectionComponent,
    ThesisComponent,
    FeedbacksuccessComponent,
    PersonalfeedbackComponent,
    UploadHeadimageComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,

    ReactiveFormsModule,
    HttpClientModule,
    FormsModule,
  ],
  providers: [GetdataService, DatePipe],
  bootstrap: [AppComponent]
})
export class AppModule { }
