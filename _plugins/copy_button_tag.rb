module Jekyll
    class CopyButtonTag < Liquid::Tag
      def render(context)
        <<~HTML
          <pre class="highlight">
            <code>
              #{super}
            </code>
            <button class="copy-button">
              <span class="tooltip">Copy</span>
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect>
                <path d="M5 15h6"></path>
                <path d="M5 9h6"></path>
                <path d="M9 5h9"></path>
              </svg>
            </button>
          </pre>
        HTML
      end
    end
  end
  
  Liquid::Template.register_tag('copybutton', Jekyll::CopyButtonTag)
  