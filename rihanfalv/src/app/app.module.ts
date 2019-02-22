import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

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
    HeadComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
