import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { FeedbackComponent } from './components/feedback/feedback.component';
import { FeedbacksuccessComponent } from './components/Feedback/feedbacksuccess/feedbacksuccess.component';
import { PersonalfeedbackComponent } from './components/Feedback/personalfeedback/personalfeedback.component';
import { LegalComponent } from './components/provisions/legal/legal.component';
import { ArticleComponent } from './components/provisions/article/article.component';
import { ContentComponent } from './components/provisions/content/content.component';

const routes: Routes = [
  {
    path:'Feedback',component:FeedbackComponent
  },
  {
    path:'Personalfeedback',component:PersonalfeedbackComponent
  },
  {
    path:'Feedbacksuccess',component:FeedbacksuccessComponent
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
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
