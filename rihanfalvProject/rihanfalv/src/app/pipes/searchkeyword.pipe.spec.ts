import { SearchkeywordPipe } from './searchkeyword.pipe';
import {DomSanitizer} from '@angular/platform-browser';

describe('SearchkeywordPipe', () => {
  it('create an instance', () => {
    let DOM:DomSanitizer
    const pipe = new SearchkeywordPipe(DOM );
    expect(pipe).toBeTruthy();
  });
});
