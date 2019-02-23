import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { FeedbackComponent } from './components/feedback/feedback.component';
import { FeedbacksuccessComponent } from './components/Feedback/feedbacksuccess/feedbacksuccess.component';
import { PersonalfeedbackComponent } from './components/Feedback/personalfeedback/personalfeedback.component';
import { LegalComponent } from './components/provisions/legal/legal.component';
import { ArticleComponent } from './components/provisions/article/article.component';
import { ContentComponent } from './components/provisions/content/content.component';
import { PaperComponent } from './components/paper/paper.component';
import { PaperwebComponent } from './components/paper/paperweb/paperweb.component';
import { CaseComponent } from './components/case/case.component';
import { CaseDataComponent } from './components/case/case-data/case-data.component';

import { PopularComponent } from '../app/components/chat/popular/popular.component';
import { PostComponent } from '../app/components/chat/post/post.component';
import { ReplyPageComponent } from '../app/components/chat/reply-page/reply-page.component';
import { ChatHeadComponent } from '../app/components/chat/chat-head/chat-head.component';
import { TagComponent } from '../app/components/chat/tag/tag.component';


const routes: Routes = [
  {
    path: 'Feedback', component: FeedbackComponent
  },
  {
    path: 'Personalfeedback', component: PersonalfeedbackComponent
  },
  {
    path: 'Feedbacksuccess', component: FeedbacksuccessComponent
  },
  {
    path: 'legal', component: LegalComponent
  },
  {
    path: 'article', component: ArticleComponent
  },
  {
    path: 'content', component: ContentComponent
  },
  {
    path: 'paper', component: PaperComponent
  },
  {
    path: 'paperweb/:ID', component: PaperwebComponent
  },
  {
    path:"CaseAnalysis",component:CaseComponent
  },
  {
    path:"display-data",component:CaseDataComponent
  },

  { path: 'popular', component: PopularComponent },
  { path: 'post', component: PostComponent },
  { path: 'replyPage', component: ReplyPageComponent },
  { path: 'head', component: ChatHeadComponent },
  { path: 'tag', component: TagComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
