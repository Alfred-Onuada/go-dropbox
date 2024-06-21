import { Input, Component } from '@angular/core';
import { MarqueeDirective } from '../../directives/marquee.directive';
import { FileSizePipe } from '../../pipes/file-size.pipe';
import IFile from '../../interfaces/file';

@Component({
  selector: 'app-video',
  standalone: true,
  imports: [MarqueeDirective, FileSizePipe],
  templateUrl: './video.component.html',
  styleUrl: './video.component.css'
})
export class VideoComponent {
  @Input() file!: IFile;
}
