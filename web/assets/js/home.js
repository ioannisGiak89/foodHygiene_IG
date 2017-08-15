"use strict";
$(document).ready(function() {

    var $form;
    var $select;
    var $alertBox;
    var $button;
    var selectedValue = "";

    function postForm() {
        $.ajax({
            type: "POST",
            url: "/ajax/api/establishments",
            data: $form.serialize(),
            dataType: "json",
            success: function (data) {
                $alertBox.addClass("hidden");
                $(".table").append(data.Body);
                enableButton();
            },
            error: function (xhr) {
                var response = xhr.responseJSON;
                showAlertBox("Something went wrong please try again!");
                enableButton()
            }
        });
    }

    function enableButton() {
        $button.text("Look up!");
        $button.attr('disabled', false);
    }

    function showAlertBox(message) {
        $alertBox.text(message);
        $alertBox.removeClass("hidden");
    }

    function attachButtonListener() {
        $button.click(function () {
            var newVal = $select.val();

            if (newVal !== "") {

                if (newVal === selectedValue) {
                    showAlertBox("Please select a new authority!");
                    return;
                }
                selectedValue = newVal;
                $("table").remove();
                $button.text("Loading..");
                $button.attr('disabled', true);
                $alertBox.addClass("hidden");
                postForm();
            } else {
                showAlertBox("Please select an authority and try again!")
            }
        });
    }

    function attachSelectListener() {
        $select.change(function(){
            $alertBox.addClass("hidden");
        });
    }

    $form = $(".home-authorities-form");
    if ($form.length > 0) {
        $select = $(".home-authorities-select");
        $alertBox = $(".alert");
        $button = $(".home-authorities-submit");

        attachButtonListener();
        attachSelectListener()
    }
});
