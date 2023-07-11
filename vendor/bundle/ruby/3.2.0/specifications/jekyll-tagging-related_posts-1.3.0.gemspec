# -*- encoding: utf-8 -*-
# stub: jekyll-tagging-related_posts 1.3.0 ruby lib

Gem::Specification.new do |s|
  s.name = "jekyll-tagging-related_posts".freeze
  s.version = "1.3.0"

  s.required_rubygems_version = Gem::Requirement.new(">= 0".freeze) if s.respond_to? :required_rubygems_version=
  s.metadata = { "changelog_uri" => "https://github.com/toshimaru/jekyll-tagging-related_posts/releases", "homepage_uri" => "https://github.com/toshimaru/jekyll-tagging-related_posts", "source_code_uri" => "https://github.com/toshimaru/jekyll-tagging-related_posts" } if s.respond_to? :metadata=
  s.require_paths = ["lib".freeze]
  s.authors = ["toshimaru".freeze]
  s.bindir = "exe".freeze
  s.date = "2023-01-29"
  s.description = "Jekyll `related_posts` function based on tags (works on Jekyll3). It replaces original Jekyll's `related_posts` function to use tags to calculate relationships.".freeze
  s.email = ["me@toshimaru.net".freeze]
  s.homepage = "https://github.com/toshimaru/jekyll-tagging-related_posts".freeze
  s.licenses = ["MIT".freeze]
  s.required_ruby_version = Gem::Requirement.new(">= 2.7.0".freeze)
  s.rubygems_version = "3.4.10".freeze
  s.summary = "Jekyll `related_posts` function based on tags (works on Jekyll3)".freeze

  s.installed_by_version = "3.4.10" if s.respond_to? :installed_by_version

  s.specification_version = 4

  s.add_runtime_dependency(%q<jekyll>.freeze, [">= 3.9", "< 5.0"])
end
