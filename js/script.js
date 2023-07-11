document.addEventListener('DOMContentLoaded', function () {
    // Get all code blocks on the page
    var codeBlocks = document.querySelectorAll('pre.highlight');
  
    // Iterate over each code block
    Array.prototype.forEach.call(codeBlocks, function (codeBlock) {
      // Create the copy button
      var copyButton = document.createElement('button');
      copyButton.className = 'copy-button';
      copyButton.setAttribute('title', 'Copy code');
      copyButton.textContent = 'Copy';
  
      // Append the copy button to the code block
      codeBlock.appendChild(copyButton);
  
      // Add click event listener to the copy button
      copyButton.addEventListener('click', function () {
        // Get the code content
        var codeContent = codeBlock.querySelector('code').innerText;
  
        // Create a textarea element to hold the code content
        var textarea = document.createElement('textarea');
        textarea.value = codeContent;
  
        // Append the textarea to the document
        document.body.appendChild(textarea);
  
        // Select the code content in the textarea
        textarea.select();
  
        // Copy the selected code to the clipboard
        document.execCommand('copy');
  
        // Remove the textarea from the document
        document.body.removeChild(textarea);
  
        // Change the button text temporarily to indicate successful copy
        copyButton.textContent = 'Copied!';
  
        // Reset the button text after a short delay
        setTimeout(function () {
          copyButton.textContent = 'Copy';
        }, 1500);
      });
    });
  });
  