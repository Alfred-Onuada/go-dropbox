import { Component, Input } from '@angular/core';
import IFile from '../../interfaces/file';
import { MarqueeDirective } from '../../directives/marquee.directive';
import { FileSizePipe } from '../../pipes/file-size.pipe';

@Component({
  selector: 'app-file',
  standalone: true,
  imports: [MarqueeDirective, FileSizePipe],
  templateUrl: './file.component.html',
  styleUrl: './file.component.css'
})
export class FileComponent {
  @Input() file!: IFile;
}
