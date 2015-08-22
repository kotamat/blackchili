$(document).ready(function(){
  var obj = $.parseJSON(json_raw);
  $.each(obj.lines, function(line_num, line_words){
    var line = '<li>';
    $.each(line_words.words, function(word_index, word){
      line += word.string;
    });
    line += '</li>';
    $('ul#lines').append(line);
  });
});
