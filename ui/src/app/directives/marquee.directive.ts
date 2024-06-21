import { AfterViewInit, Directive, ElementRef, Renderer2 } from '@angular/core';

@Directive({
  selector: '[appMarquee]',
  standalone: true
})
export class MarqueeDirective implements AfterViewInit {
  constructor(private el: ElementRef, private renderer: Renderer2) {}

  ngAfterViewInit() {
    const container = this.el.nativeElement;
    const content = container.firstChild;

    // Add styles for container and content
    this.renderer.setStyle(container, 'overflow', 'hidden');
    this.renderer.setStyle(container, 'white-space', 'nowrap');
    this.renderer.setStyle(container, 'display', 'inline-block');
    
    // Check if content width exceeds container width
    if (content.scrollWidth > container.clientWidth) {
      this.renderer.setStyle(content, 'display', 'inline-block');
      this.renderer.setStyle(content, 'animation', 'marquee 5s linear infinite');
    }
  }
}