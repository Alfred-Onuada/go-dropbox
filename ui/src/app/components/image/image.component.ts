import { Component, Input } from '@angular/core';
import IFile from '../../interfaces/file';
import { MarqueeDirective } from '../../directives/marquee.directive';
import { FileSizePipe } from '../../pipes/file-size.pipe';

@Component({
  selector: 'app-image',
  standalone: true,
  imports: [MarqueeDirective, FileSizePipe],
  templateUrl: './image.component.html',
  styleUrl: './image.component.css'
})
export class ImageComponent {
  @Input() file!: IFile;
}
