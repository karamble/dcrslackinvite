
$(window).load(function () {

	/* for css cache
	$('link').each( function() {
		$(this).attr('href', $(this).attr('href')+'?'+Math.random().toString(36).substr(2, 5));
	});*/

	// for json API
	$.ajaxSetup({
		async: false
	});

		// font weight
		fontRegular = 'fontregular',
		fontSemibold = 'fontsemibold',
		fontBold = 'fontbold',

		// font size
		font14 = 'font14',
		font16 = 'font16',
		font18 = 'font18',
		font22 = 'font22',
		font24 = 'font24',
		font28 = 'font28',
		font38 = 'font38',

		// font color
		colorDarkBlue = 'colordarkblue',

		// line color
		verticalBlue = 'verticalblue',

		// bg color
		backgroundDarkBlue = 'backgrounddarkblue',
		backgroundCyan = 'backgroundcyan',
		backgroundGray = 'backgroundgray',
		backgroundBlue = 'backgroundblue',

		// solid color
		turquoise = 'turquoise';

		// transition
		transition = 'transition',
		transitionModest = 'transitionmodest',
		transitionSlow = 'transitionslow',

		// addins
		guideBlockContentLast = 'guideblockcontentlast',
		cursor = 'cursor',
		hand = 'hand',
		active = 'active',
		counter = 1,

		// footer
		footerBlock = $('.footerblock'),
		icon = $('.icon'),
		footerBlockIndicator = $('.footerblockindicator'),

		// navigation
		navigationButton = $('.navigation-button'),
		navOpenClose = $('.nav-open-close'),
		linkSection = $('.link-section'),

	// navigation menu
	navigationButton.click( function() {
		navOpenClose.add(linkSection).toggleClass('active');
	});
});


