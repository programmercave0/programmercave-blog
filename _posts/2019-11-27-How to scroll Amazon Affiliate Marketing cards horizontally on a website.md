---
layout: post
title: "How to scroll Amazon Affiliate Marketing cards horizontally on a website"
author: "Programmercave"
tags:  [HowTo, HTML, CSS]
date: 2019-11-27
---



I wanted to put my amazon affiliate marketing product on website. But on smaller screens it was messing up my website's interface. I searched on internet tried many things but no result. Then I tried it with `<table>` in html and css and it worked. Here is the code for that.

Table with only one reow is needed and we will put products' card on that row and when the table width is greater than screen's width it will scroll horizontaly.

```
<div style="overflow-x:auto;">
 <table>
    <tr>
<!-- Acer Aspire E15 -->
	<th><iframe name ="card" style="width:120px;height:240px;" marginwidth="0" marginheight="0" scrolling="no" frameborder="0" src="//ws-in.amazon-adsystem.com/widgets/q?ServiceVersion=20070822&OneJS=1&Operation=GetAdHtml&MarketPlace=IN&source=ac&ref=qf_sp_asin_til&ad_type=product_link&tracking_id=0106d-21&marketplace=amazon&region=IN&placement=B07MK7G18V&asins=B07MK7G18V&linkId=2f4d4031c2429525b56637f2704b59ec&show_border=true&link_opens_in_new_window=true&price_color=333333&title_color=0066c0&bg_color=ffffff">
    </iframe></th>
    
 <!-- Asus Vivobook -->
 	<th><iframe name="card" style="width:120px;height:240px;" marginwidth="0" marginheight="0" scrolling="no" frameborder="0" src="//ws-in.amazon-adsystem.com/widgets/q?ServiceVersion=20070822&OneJS=1&Operation=GetAdHtml&MarketPlace=IN&source=ac&ref=qf_sp_asin_til&ad_type=product_link&tracking_id=0106d-21&marketplace=amazon&region=IN&placement=B07N8CQN2F&asins=B07N8CQN2F&linkId=587e7b91213538f4435f5376bd0e7fa1&show_border=true&link_opens_in_new_window=true&price_color=333333&title_color=0066c0&bg_color=ffffff">
    </iframe></th>
    
 <!-- Lenovo Thinkpad-->
 	<th><iframe name="card" style="width:120px;height:240px;" marginwidth="0" marginheight="0" scrolling="no" frameborder="0" src="//ws-in.amazon-adsystem.com/widgets/q?ServiceVersion=20070822&OneJS=1&Operation=GetAdHtml&MarketPlace=IN&source=ac&ref=qf_sp_asin_til&ad_type=product_link&tracking_id=0106d-21&marketplace=amazon&region=IN&placement=B07TCV9VB8&asins=B07TCV9VB8&linkId=0110e8a0b000b86dea08d41c80a1b9ec&show_border=true&link_opens_in_new_window=true&price_color=333333&title_color=0066c0&bg_color=ffffff">
    </iframe></th>
    
 <!-- Dell G5 series-->
 	<th><iframe name="card" style="width:120px;height:240px;" marginwidth="0" marginheight="0" scrolling="no" frameborder="0" src="//ws-in.amazon-adsystem.com/widgets/q?ServiceVersion=20070822&OneJS=1&Operation=GetAdHtml&MarketPlace=IN&source=ac&ref=qf_sp_asin_til&ad_type=product_link&tracking_id=0106d-21&marketplace=amazon&region=IN&placement=B07B44TMBF&asins=B07B44TMBF&linkId=b11de009d21f7476ad34675e69e6fb58&show_border=true&link_opens_in_new_window=true&price_color=333333&title_color=0066c0&bg_color=ffffff">
    </iframe></th>

<!-- HP 15 -->
	<th><iframe name="card" style="width:120px;height:240px;" marginwidth="0" marginheight="0" scrolling="no" frameborder="0" src="//ws-in.amazon-adsystem.com/widgets/q?ServiceVersion=20070822&OneJS=1&Operation=GetAdHtml&MarketPlace=IN&source=ac&ref=qf_sp_asin_til&ad_type=product_link&tracking_id=0106d-21&marketplace=amazon&region=IN&placement=B07HHS5TZT&asins=B07HHS5TZT&linkId=b8f4ae0b5d3cd85a3789c668c121b1a3&show_border=true&link_opens_in_new_window=true&price_color=333333&title_color=0066c0&bg_color=ffffff">
    </iframe></th>

<!-- Asus k50 -->
	<th><iframe name="card" style="width:120px;height:240px;" marginwidth="0" marginheight="0" scrolling="no" frameborder="0" src="//ws-in.amazon-adsystem.com/widgets/q?ServiceVersion=20070822&OneJS=1&Operation=GetAdHtml&MarketPlace=IN&source=ac&ref=qf_sp_asin_til&ad_type=product_link&tracking_id=0106d-21&marketplace=amazon&region=IN&placement=B07RRQBC1S&asins=B07RRQBC1S&linkId=5d55270eceda8ed147ca9d0bb15c73f6&show_border=true&link_opens_in_new_window=true&price_color=333333&title_color=0066c0&bg_color=ffffff">
    </iframe></th>

<!--Dell XPS-->
	<th><iframe name="card" style="width:120px;height:240px;" marginwidth="0" marginheight="0" scrolling="no" frameborder="0" src="//ws-in.amazon-adsystem.com/widgets/q?ServiceVersion=20070822&OneJS=1&Operation=GetAdHtml&MarketPlace=IN&source=ac&ref=qf_sp_asin_til&ad_type=product_link&tracking_id=0106d-21&marketplace=amazon&region=IN&placement=B07B2X7C5C&asins=B07B2X7C5C&linkId=c132a6e522b50ca7158785254eb2c283&show_border=true&link_opens_in_new_window=true&price_color=333333&title_color=0066c0&bg_color=ffffff">
    	</iframe></th>
     </tr>
  </table>
</div>
```

Now to customize table the css code.

{% include ads.html %}<br/>

```
table {
  border-spacing: 0;
  border-collapse: collapse;
  width: 100%;
  border: 1px solid #ddd;
}
td,
th {
  text-align: left;
  padding: 3px;
}
```

You can find these codes in this repository.

{% include ads.html %}<br/>

This is how it will look.

{% include ads.html %}<br/>

