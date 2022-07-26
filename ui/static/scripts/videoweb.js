$(document).ready(function() {
    $('#videos a').click(function() {
       var data = $(this).attr('data');
       $('#loader').append('<video src="assets/videos/'+data+'" controls></video>');
       $('#overlay').fadeIn(250);
    });
 
    $('#close').click(function() {
       $('#overlay').fadeOut(250,function() {
          $('#loader').html('');
       });
    });
 });