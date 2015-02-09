var AutoreplyText = function() {


  return {
    //main function to initiate the module
    init: function() {
      //text
      $("#textsave").click(function() {
        $content = $("#textcontent").val();

        $keyword = $("#textkeyword").val();
        $flag = $("#textflag").val();
        $.get("http://121.40.35.3/mongoinsertkeyword?keyword=" + $keyword + "&content=" + $content + "&flag=" + $flag,
          function(data, status) {
            // alert(data);
            window.location.href = "/autoreply";
          });
      });
      $("#textcancle").click(function() {

        window.location.href = "/autoreply";

      });



      //  news
      $("#newstitle").bind('input propertychange', function() {

        $('.newstitleshow').html($("#newstitle").val());
      });

      $("#newsdescribe").bind('input propertychange', function() {

        $('.newsdescribeshow').html($("#newsdescribe").val());
      });

      $("#newssave").click(function() {


        $newskeyword = $("#newskeyword").val();
        $newsflag = $("#newsflag").val();
        $newstitle = $("#newstitle").val();

        $newsauth = $("#newsauth").val();
        $newsdescribe = $("#newsdescribe").val();

        // $.get("http://121.40.35.3/mongoinsertkeyword?keyword="+$keyword+"&content="+$content+"&flag="+$flag,
        //     function(data,status){
        //    // alert(data);
        //       window.location.href = "/autoreply";
        //  });
      });
      $("#newscancle").click(function() {

        window.location.href = "/autoreply";

      });
      //news  

      //file
      $('#fileupload').fileupload({
        dataType: 'json',
        done: function(e, data) {
          $.each(data.result.files, function(index, file) {
            $('<p/>').text(file.name).appendTo(document.body);
          });
        }
      });
      //file

    }

  };

}();