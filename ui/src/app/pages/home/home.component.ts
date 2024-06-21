import { Component, Signal, signal } from '@angular/core';
import IFile from '../../interfaces/file';
import { AudioComponent } from '../../components/audio/audio.component';
import { FileComponent } from '../../components/file/file.component';
import { VideoComponent } from '../../components/video/video.component';
import { ImageComponent } from '../../components/image/image.component';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [AudioComponent, FileComponent, VideoComponent, ImageComponent],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent {
  files: Signal<IFile[]> = signal([
    {
      name: "Beethoven_Symphony_No5",
      size: 5000000,
      extension: "mp3",
      type: "audio/mpeg"
    },
    {
      name: "Nature_Documentary",
      size: 20000000,
      extension: "mp4",
      type: "video/mp4"
    },
    {
      name: "Sunset_in_Hawaii",
      size: 3000000,
      extension: "jpg",
      type: "image/jpeg"
    },
    {
      name: "Project_Report",
      size: 150000,
      extension: "pdf",
      type: "application/pdf"
    },
    {
      name: "Jazz_Classics",
      size: 4500000,
      extension: "wav",
      type: "audio/wav"
    },
    {
      name: "Holiday_Trip",
      size: 25000000,
      extension: "mkv",
      type: "video/x-matroska"
    },
    {
      name: "Family_Reunion",
      size: 4000000,
      extension: "png",
      type: "image/png"
    },
    {
      name: "Thesis_Document",
      size: 100000,
      extension: "docx",
      type: "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
    },
    {
      name: "Classical_Piano_Collection",
      size: 5200000,
      extension: "flac",
      type: "audio/flac"
    },
    {
      name: "Wildlife_Adventure",
      size: 18000000,
      extension: "avi",
      type: "video/x-msvideo"
    },
    {
      name: "Beethoven_Symphony_No5",
      size: 5000000,
      extension: "mp3",
      type: "audio/mpeg"
    },
    {
      name: "Nature_Documentary",
      size: 20000000,
      extension: "mp4",
      type: "video/mp4"
    },
    {
      name: "Sunset_in_Hawaii",
      size: 3000000,
      extension: "jpg",
      type: "image/jpeg"
    },
    {
      name: "Project_Report",
      size: 150000,
      extension: "pdf",
      type: "application/pdf"
    },
    {
      name: "Jazz_Classics",
      size: 4500000,
      extension: "wav",
      type: "audio/wav"
    },
    {
      name: "Holiday_Trip",
      size: 25000000,
      extension: "mkv",
      type: "video/x-matroska"
    },
    {
      name: "Family_Reunion",
      size: 4000000,
      extension: "png",
      type: "image/png"
    },
    {
      name: "Thesis_Document",
      size: 100000,
      extension: "docx",
      type: "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
    },
    {
      name: "Classical_Piano_Collection",
      size: 5200000,
      extension: "flac",
      type: "audio/flac"
    },
    {
      name: "Wildlife_Adventure",
      size: 18000000,
      extension: "avi",
      type: "video/x-msvideo"
    },
    {
      name: "Beethoven_Symphony_No5",
      size: 5000000,
      extension: "mp3",
      type: "audio/mpeg"
    },
    {
      name: "Nature_Documentary",
      size: 20000000,
      extension: "mp4",
      type: "video/mp4"
    },
    {
      name: "Sunset_in_Hawaii",
      size: 3000000,
      extension: "jpg",
      type: "image/jpeg"
    },
    {
      name: "Project_Report",
      size: 150000,
      extension: "pdf",
      type: "application/pdf"
    },
    {
      name: "Jazz_Classics",
      size: 4500000,
      extension: "wav",
      type: "audio/wav"
    },
    {
      name: "Holiday_Trip",
      size: 25000000,
      extension: "mkv",
      type: "video/x-matroska"
    },
    {
      name: "Family_Reunion",
      size: 4000000,
      extension: "png",
      type: "image/png"
    },
    {
      name: "Thesis_Document",
      size: 100000,
      extension: "docx",
      type: "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
    },
    {
      name: "Classical_Piano_Collection",
      size: 5200000,
      extension: "flac",
      type: "audio/flac"
    },
    {
      name: "Wildlife_Adventure",
      size: 18000000,
      extension: "avi",
      type: "video/x-msvideo"
    },
    {
      name: "Beethoven_Symphony_No5",
      size: 5000000,
      extension: "mp3",
      type: "audio/mpeg"
    },
    {
      name: "Nature_Documentary",
      size: 20000000,
      extension: "mp4",
      type: "video/mp4"
    },
    {
      name: "Sunset_in_Hawaii",
      size: 3000000,
      extension: "jpg",
      type: "image/jpeg"
    },
    {
      name: "Project_Report",
      size: 150000,
      extension: "pdf",
      type: "application/pdf"
    },
    {
      name: "Jazz_Classics",
      size: 4500000,
      extension: "wav",
      type: "audio/wav"
    },
    {
      name: "Holiday_Trip",
      size: 25000000,
      extension: "mkv",
      type: "video/x-matroska"
    },
    {
      name: "Family_Reunion",
      size: 4000000,
      extension: "png",
      type: "image/png"
    },
    {
      name: "Thesis_Document",
      size: 100000,
      extension: "docx",
      type: "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
    },
    {
      name: "Classical_Piano_Collection",
      size: 5200000,
      extension: "flac",
      type: "audio/flac"
    },
    {
      name: "Wildlife_Adventure",
      size: 18000000,
      extension: "avi",
      type: "video/x-msvideo"
    },
    {
      name: "Beethoven_Symphony_No5",
      size: 5000000,
      extension: "mp3",
      type: "audio/mpeg"
    },
    {
      name: "Nature_Documentary",
      size: 20000000,
      extension: "mp4",
      type: "video/mp4"
    },
    {
      name: "Sunset_in_Hawaii",
      size: 3000000,
      extension: "jpg",
      type: "image/jpeg"
    },
    {
      name: "Project_Report",
      size: 150000,
      extension: "pdf",
      type: "application/pdf"
    },
    {
      name: "Jazz_Classics",
      size: 4500000,
      extension: "wav",
      type: "audio/wav"
    },
    {
      name: "Holiday_Trip",
      size: 25000000,
      extension: "mkv",
      type: "video/x-matroska"
    },
    {
      name: "Family_Reunion",
      size: 4000000,
      extension: "png",
      type: "image/png"
    },
    {
      name: "Thesis_Document",
      size: 100000,
      extension: "docx",
      type: "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
    },
    {
      name: "Classical_Piano_Collection",
      size: 5200000,
      extension: "flac",
      type: "audio/flac"
    },
    {
      name: "Wildlife_Adventure",
      size: 18000000,
      extension: "avi",
      type: "video/x-msvideo"
    },
    {
      name: "Beethoven_Symphony_No5",
      size: 5000000,
      extension: "mp3",
      type: "audio/mpeg"
    },
    {
      name: "Nature_Documentary",
      size: 20000000,
      extension: "mp4",
      type: "video/mp4"
    },
    {
      name: "Sunset_in_Hawaii",
      size: 3000000,
      extension: "jpg",
      type: "image/jpeg"
    },
    {
      name: "Project_Report",
      size: 150000,
      extension: "pdf",
      type: "application/pdf"
    },
    {
      name: "Jazz_Classics",
      size: 4500000,
      extension: "wav",
      type: "audio/wav"
    },
    {
      name: "Holiday_Trip",
      size: 25000000,
      extension: "mkv",
      type: "video/x-matroska"
    },
    {
      name: "Family_Reunion",
      size: 4000000,
      extension: "png",
      type: "image/png"
    },
    {
      name: "Thesis_Document",
      size: 100000,
      extension: "docx",
      type: "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
    },
    {
      name: "Classical_Piano_Collection",
      size: 5200000,
      extension: "flac",
      type: "audio/flac"
    },
    {
      name: "Wildlife_Adventure",
      size: 18000000,
      extension: "avi",
      type: "video/x-msvideo"
    },
    {
      name: "Beethoven_Symphony_No5",
      size: 5000000,
      extension: "mp3",
      type: "audio/mpeg"
    },
    {
      name: "Nature_Documentary",
      size: 20000000,
      extension: "mp4",
      type: "video/mp4"
    },
    {
      name: "Sunset_in_Hawaii",
      size: 3000000,
      extension: "jpg",
      type: "image/jpeg"
    },
    {
      name: "Project_Report",
      size: 150000,
      extension: "pdf",
      type: "application/pdf"
    },
    {
      name: "Jazz_Classics",
      size: 4500000,
      extension: "wav",
      type: "audio/wav"
    },
    {
      name: "Holiday_Trip",
      size: 25000000,
      extension: "mkv",
      type: "video/x-matroska"
    },
    {
      name: "Family_Reunion",
      size: 4000000,
      extension: "png",
      type: "image/png"
    },
    {
      name: "Thesis_Document",
      size: 100000,
      extension: "docx",
      type: "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
    },
    {
      name: "Classical_Piano_Collection",
      size: 5200000,
      extension: "flac",
      type: "audio/flac"
    },
    {
      name: "Wildlife_Adventure",
      size: 18000000,
      extension: "avi",
      type: "video/x-msvideo"
    }
  ]);

  currdir: Signal<string> = signal('/Users/alfredonuada/Documents/portfolio/fun_and_play/go-dropbox/server');
}
