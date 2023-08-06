module Jekyll
    module TOCGenerator
      def toc(content)
        toc = ""
        headers = []
        content.scan(/<h(\d)>(.*?)<\/h\d>/i) do |level, title|
          headers << { level: level.to_i, title: title.strip }
        end
  
        headers.each_with_index do |header, index|
          indent = "  " * (header[:level] - 1)
          toc << "#{indent}- [#{header[:title]}](#header#{index + 1})\n"
        end
  
        toc
      end
    end
  end
  
  Liquid::Template.register_filter(Jekyll::TOCGenerator)
  