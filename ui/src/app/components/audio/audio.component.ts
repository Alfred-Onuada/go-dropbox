import { Component, Input } from '@angular/core';
import IFile from '../../interfaces/file';
import { MarqueeDirective } from '../../directives/marquee.directive';
import { FileSizePipe } from '../../pipes/file-size.pipe';

@Component({
  selector: 'app-audio',
  standalone: true,
  imports: [MarqueeDirective, FileSizePipe],
  templateUrl: './audio.component.html',
  styleUrl: './audio.component.css'
})
export class AudioComponent {
  @Input() file!: IFile;
}
