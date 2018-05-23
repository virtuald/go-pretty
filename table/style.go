package table

import (
	"github.com/jedib0t/go-pretty/text"
)

// Style declares how to render the Table and provides very fine-grained control
// on how the Table gets rendered on the Console.
type Style struct {
	Name    string
	Box     BoxStyle
	Color   ColorOptions
	Format  FormatOptions
	Options Options
}

var (
	// StyleDefault renders a Table like below:
	//  +-----+------------+-----------+--------+-----------------------------+
	//  |   # | FIRST NAME | LAST NAME | SALARY |                             |
	//  +-----+------------+-----------+--------+-----------------------------+
	//  |   1 | Arya       | Stark     |   3000 |                             |
	//  |  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |
	//  | 300 | Tyrion     | Lannister |   5000 |                             |
	//  +-----+------------+-----------+--------+-----------------------------+
	//  |     |            | TOTAL     |  10000 |                             |
	//  +-----+------------+-----------+--------+-----------------------------+
	StyleDefault = Style{
		Name:    "StyleDefault",
		Box:     StyleBoxDefault,
		Color:   ColorOptionsDefault,
		Format:  FormatOptionsDefault,
		Options: OptionsDefault,
	}

	// StyleBold renders a Table like below:
	//  ┏━━━━━┳━━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
	//  ┃   # ┃ FIRST NAME ┃ LAST NAME ┃ SALARY ┃                             ┃
	//  ┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
	//  ┃   1 ┃ Arya       ┃ Stark     ┃   3000 ┃                             ┃
	//  ┃  20 ┃ Jon        ┃ Snow      ┃   2000 ┃ You know nothing, Jon Snow! ┃
	//  ┃ 300 ┃ Tyrion     ┃ Lannister ┃   5000 ┃                             ┃
	//  ┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
	//  ┃     ┃            ┃ TOTAL     ┃  10000 ┃                             ┃
	//  ┗━━━━━┻━━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━┻━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
	StyleBold = Style{
		Name:    "StyleBold",
		Box:     StyleBoxBold,
		Color:   ColorOptionsDefault,
		Format:  FormatOptionsDefault,
		Options: OptionsDefault,
	}

	// StyleColoredBright renders a Table without any borders or separators,
	// and with Black text on Cyan background for Header/Footer and
	// White background for other rows.
	StyleColoredBright = Style{
		Name:    "StyleColoredBright",
		Box:     StyleBoxDefault,
		Color:   ColorOptionsBright,
		Format:  FormatOptionsDefault,
		Options: OptionsNoBordersAndSeparators,
	}

	// StyleColoredDark renders a Table without any borders or separators, and
	// with Header/Footer in Cyan text and other rows with White text, all on
	// Black background.
	StyleColoredDark = Style{
		Name:    "StyleColoredDark",
		Box:     StyleBoxDefault,
		Color:   ColorOptionsDark,
		Format:  FormatOptionsDefault,
		Options: OptionsNoBordersAndSeparators,
	}

	// StyleColoredBlackOnBlueWhite renders a Table without any borders or
	// separators, and with Black text on Blue background for Header/Footer and
	// White background for other rows.
	StyleColoredBlackOnBlueWhite = Style{
		Name:    "StyleColoredBlackOnBlueWhite",
		Box:     StyleBoxDefault,
		Color:   ColorOptionsBlackOnBlueWhite,
		Format:  FormatOptionsDefault,
		Options: OptionsNoBordersAndSeparators,
	}

	// StyleColoredBlackOnCyanWhite renders a Table without any borders or
	// separators, and with Black text on Cyan background for Header/Footer and
	// White background for other rows.
	StyleColoredBlackOnCyanWhite = Style{
		Name:    "StyleColoredBlackOnCyanWhite",
		Box:     StyleBoxDefault,
		Color:   ColorOptionsBlackOnCyanWhite,
		Format:  FormatOptionsDefault,
		Options: OptionsNoBordersAndSeparators,
	}

	// StyleColoredBlackOnGreenWhite renders a Table without any borders or
	// separators, and with Black text on Green background for Header/Footer and
	// White background for other rows.
	StyleColoredBlackOnGreenWhite = Style{
		Name:    "StyleColoredBlackOnGreenWhite",
		Box:     StyleBoxDefault,
		Color:   ColorOptionsBlackOnGreenWhite,
		Format:  FormatOptionsDefault,
		Options: OptionsNoBordersAndSeparators,
	}

	// StyleColoredBlackOnMagentaWhite renders a Table without any borders or
	// separators, and with Black text on Magenta background for Header/Footer and
	// White background for other rows.
	StyleColoredBlackOnMagentaWhite = Style{
		Name:    "StyleColoredBlackOnMagentaWhite",
		Box:     StyleBoxDefault,
		Color:   ColorOptionsBlackOnMagentaWhite,
		Format:  FormatOptionsDefault,
		Options: OptionsNoBordersAndSeparators,
	}

	// StyleColoredBlackOnYellowWhite renders a Table without any borders or
	// separators, and with Black text on Yellow background for Header/Footer and
	// White background for other rows.
	StyleColoredBlackOnYellowWhite = Style{
		Name:    "StyleColoredBlackOnYellowWhite",
		Box:     StyleBoxDefault,
		Color:   ColorOptionsBlackOnYellowWhite,
		Format:  FormatOptionsDefault,
		Options: OptionsNoBordersAndSeparators,
	}

	// StyleColoredBlackOnRedWhite renders a Table without any borders or
	// separators, and with Black text on Red background for Header/Footer and
	// White background for other rows.
	StyleColoredBlackOnRedWhite = Style{
		Name:    "StyleColoredBlackOnRedWhite",
		Box:     StyleBoxDefault,
		Color:   ColorOptionsBlackOnRedWhite,
		Format:  FormatOptionsDefault,
		Options: OptionsNoBordersAndSeparators,
	}

	// StyleColoredBlueWhiteOnBlack renders a Table without any borders or
	// separators, and with Header/Footer in Blue text and other rows with
	// White text, all on Black background.
	StyleColoredBlueWhiteOnBlack = Style{
		Name:    "StyleColoredBlueWhiteOnBlack",
		Box:     StyleBoxDefault,
		Color:   ColorOptionsBlueWhiteOnBlack,
		Format:  FormatOptionsDefault,
		Options: OptionsNoBordersAndSeparators,
	}

	// StyleColoredCyanWhiteOnBlack renders a Table without any borders or
	// separators, and with Header/Footer in Cyan text and other rows with
	// White text, all on Black background.
	StyleColoredCyanWhiteOnBlack = Style{
		Name:    "StyleColoredCyanWhiteOnBlack",
		Box:     StyleBoxDefault,
		Color:   ColorOptionsCyanWhiteOnBlack,
		Format:  FormatOptionsDefault,
		Options: OptionsNoBordersAndSeparators,
	}

	// StyleColoredGreenWhiteOnBlack renders a Table without any borders or
	// separators, and with Header/Footer in Green text and other rows with
	// White text, all on Black background.
	StyleColoredGreenWhiteOnBlack = Style{
		Name:    "StyleColoredGreenWhiteOnBlack",
		Box:     StyleBoxDefault,
		Color:   ColorOptionsGreenWhiteOnBlack,
		Format:  FormatOptionsDefault,
		Options: OptionsNoBordersAndSeparators,
	}

	// StyleColoredMagentaWhiteOnBlack renders a Table without any borders or
	// separators, and with Header/Footer in Magenta text and other rows with
	// White text, all on Black background.
	StyleColoredMagentaWhiteOnBlack = Style{
		Name:    "StyleColoredMagentaWhiteOnBlack",
		Box:     StyleBoxDefault,
		Color:   ColorOptionsMagentaWhiteOnBlack,
		Format:  FormatOptionsDefault,
		Options: OptionsNoBordersAndSeparators,
	}

	// StyleColoredRedWhiteOnBlack renders a Table without any borders or
	// separators, and with Header/Footer in Red text and other rows with
	// White text, all on Black background.
	StyleColoredRedWhiteOnBlack = Style{
		Name:    "StyleColoredRedWhiteOnBlack",
		Box:     StyleBoxDefault,
		Color:   ColorOptionsRedWhiteOnBlack,
		Format:  FormatOptionsDefault,
		Options: OptionsNoBordersAndSeparators,
	}

	// StyleColoredYellowWhiteOnBlack renders a Table without any borders or
	// separators, and with Header/Footer in Yellow text and other rows with
	// White text, all on Black background.
	StyleColoredYellowWhiteOnBlack = Style{
		Name:    "StyleColoredYellowWhiteOnBlack",
		Box:     StyleBoxDefault,
		Color:   ColorOptionsYellowWhiteOnBlack,
		Format:  FormatOptionsDefault,
		Options: OptionsNoBordersAndSeparators,
	}

	// StyleDouble renders a Table like below:
	//  ╔═════╦════════════╦═══════════╦════════╦═════════════════════════════╗
	//  ║   # ║ FIRST NAME ║ LAST NAME ║ SALARY ║                             ║
	//  ╠═════╬════════════╬═══════════╬════════╬═════════════════════════════╣
	//  ║   1 ║ Arya       ║ Stark     ║   3000 ║                             ║
	//  ║  20 ║ Jon        ║ Snow      ║   2000 ║ You know nothing, Jon Snow! ║
	//  ║ 300 ║ Tyrion     ║ Lannister ║   5000 ║                             ║
	//  ╠═════╬════════════╬═══════════╬════════╬═════════════════════════════╣
	//  ║     ║            ║ TOTAL     ║  10000 ║                             ║
	//  ╚═════╩════════════╩═══════════╩════════╩═════════════════════════════╝
	StyleDouble = Style{
		Name:    "StyleDouble",
		Box:     StyleBoxDouble,
		Color:   ColorOptionsDefault,
		Format:  FormatOptionsDefault,
		Options: OptionsDefault,
	}

	// StyleLight renders a Table like below:
	//  ┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
	//  │   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │   1 │ Arya       │ Stark     │   3000 │                             │
	//  │  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
	//  │ 300 │ Tyrion     │ Lannister │   5000 │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │     │            │ TOTAL     │  10000 │                             │
	//  └─────┴────────────┴───────────┴────────┴─────────────────────────────┘
	StyleLight = Style{
		Name:    "StyleLight",
		Box:     StyleBoxLight,
		Color:   ColorOptionsDefault,
		Format:  FormatOptionsDefault,
		Options: OptionsDefault,
	}

	// StyleRounded renders a Table like below:
	//  ╭─────┬────────────┬───────────┬────────┬─────────────────────────────╮
	//  │   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │   1 │ Arya       │ Stark     │   3000 │                             │
	//  │  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
	//  │ 300 │ Tyrion     │ Lannister │   5000 │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │     │            │ TOTAL     │  10000 │                             │
	//  ╰─────┴────────────┴───────────┴────────┴─────────────────────────────╯
	StyleRounded = Style{
		Name:    "StyleRounded",
		Box:     StyleBoxRounded,
		Color:   ColorOptionsDefault,
		Format:  FormatOptionsDefault,
		Options: OptionsDefault,
	}

	// styleTest renders a Table like below:
	//  (-----^------------^-----------^--------^-----------------------------)
	//  [<  #>|<FIRST NAME>|<LAST NAME>|<SALARY>|<                           >]
	//  {-----+------------+-----------+--------+-----------------------------}
	//  [<  1>|<Arya      >|<Stark    >|<  3000>|<                           >]
	//  [< 20>|<Jon       >|<Snow     >|<  2000>|<You know nothing, Jon Snow!>]
	//  [<300>|<Tyrion    >|<Lannister>|<  5000>|<                           >]
	//  {-----+------------+-----------+--------+-----------------------------}
	//  [<   >|<          >|<TOTAL    >|< 10000>|<                           >]
	//  \-----v------------v-----------v--------v-----------------------------/
	styleTest = Style{
		Name:    "styleTest",
		Box:     styleBoxTest,
		Color:   ColorOptionsDefault,
		Format:  FormatOptionsDefault,
		Options: OptionsDefault,
	}
)

// BoxStyle defines the characters/strings to use to render the borders and
// separators for the Table.
type BoxStyle struct {
	BottomLeft       string
	BottomRight      string
	BottomSeparator  string
	Left             string
	LeftSeparator    string
	MiddleHorizontal string
	MiddleSeparator  string
	MiddleVertical   string
	PaddingLeft      string
	PaddingRight     string
	PageSeparator    string
	Right            string
	RightSeparator   string
	TopLeft          string
	TopRight         string
	TopSeparator     string
	UnfinishedRow    string
}

var (
	// StyleBoxDefault defines a Boxed-Table like below:
	//  +-----+------------+-----------+--------+-----------------------------+
	//  |   # | FIRST NAME | LAST NAME | SALARY |                             |
	//  +-----+------------+-----------+--------+-----------------------------+
	//  |   1 | Arya       | Stark     |   3000 |                             |
	//  |  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |
	//  | 300 | Tyrion     | Lannister |   5000 |                             |
	//  +-----+------------+-----------+--------+-----------------------------+
	//  |     |            | TOTAL     |  10000 |                             |
	//  +-----+------------+-----------+--------+-----------------------------+
	StyleBoxDefault = BoxStyle{
		BottomLeft:       "+",
		BottomRight:      "+",
		BottomSeparator:  "+",
		Left:             "|",
		LeftSeparator:    "+",
		MiddleHorizontal: "-",
		MiddleSeparator:  "+",
		MiddleVertical:   "|",
		PaddingLeft:      " ",
		PaddingRight:     " ",
		PageSeparator:    "\n",
		Right:            "|",
		RightSeparator:   "+",
		TopLeft:          "+",
		TopRight:         "+",
		TopSeparator:     "+",
		UnfinishedRow:    " ~",
	}

	// StyleBoxBold defines a Boxed-Table like below:
	//  ┏━━━━━┳━━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
	//  ┃   # ┃ FIRST NAME ┃ LAST NAME ┃ SALARY ┃                             ┃
	//  ┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
	//  ┃   1 ┃ Arya       ┃ Stark     ┃   3000 ┃                             ┃
	//  ┃  20 ┃ Jon        ┃ Snow      ┃   2000 ┃ You know nothing, Jon Snow! ┃
	//  ┃ 300 ┃ Tyrion     ┃ Lannister ┃   5000 ┃                             ┃
	//  ┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
	//  ┃     ┃            ┃ TOTAL     ┃  10000 ┃                             ┃
	//  ┗━━━━━┻━━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━┻━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
	StyleBoxBold = BoxStyle{
		BottomLeft:       text.BoxBottomLeftBold,
		BottomRight:      text.BoxBottomRightBold,
		BottomSeparator:  text.BoxBottomSeparatorBold,
		Left:             text.BoxLeftBold,
		LeftSeparator:    text.BoxLeftSeparatorBold,
		MiddleHorizontal: text.BoxHorizontalBold,
		MiddleSeparator:  text.BoxSeparatorBold,
		MiddleVertical:   text.BoxVerticalBold,
		PaddingLeft:      " ",
		PaddingRight:     " ",
		PageSeparator:    "\n",
		Right:            text.BoxRightBold,
		RightSeparator:   text.BoxRightSeparatorBold,
		TopLeft:          text.BoxTopLeftBold,
		TopRight:         text.BoxTopRightBold,
		TopSeparator:     text.BoxTopSeparatorBold,
		UnfinishedRow:    " " + text.BoxUnfinishedLine,
	}

	// StyleBoxDouble defines a Boxed-Table like below:
	//  ╔═════╦════════════╦═══════════╦════════╦═════════════════════════════╗
	//  ║   # ║ FIRST NAME ║ LAST NAME ║ SALARY ║                             ║
	//  ╠═════╬════════════╬═══════════╬════════╬═════════════════════════════╣
	//  ║   1 ║ Arya       ║ Stark     ║   3000 ║                             ║
	//  ║  20 ║ Jon        ║ Snow      ║   2000 ║ You know nothing, Jon Snow! ║
	//  ║ 300 ║ Tyrion     ║ Lannister ║   5000 ║                             ║
	//  ╠═════╬════════════╬═══════════╬════════╬═════════════════════════════╣
	//  ║     ║            ║ TOTAL     ║  10000 ║                             ║
	//  ╚═════╩════════════╩═══════════╩════════╩═════════════════════════════╝
	StyleBoxDouble = BoxStyle{
		BottomLeft:       text.BoxBottomLeftDouble,
		BottomRight:      text.BoxBottomRightDouble,
		BottomSeparator:  text.BoxBottomSeparatorDouble,
		Left:             text.BoxLeftDouble,
		LeftSeparator:    text.BoxLeftSeparatorDouble,
		MiddleHorizontal: text.BoxHorizontalDouble,
		MiddleSeparator:  text.BoxSeparatorDouble,
		MiddleVertical:   text.BoxVerticalDouble,
		PaddingLeft:      " ",
		PaddingRight:     " ",
		PageSeparator:    "\n",
		Right:            text.BoxRightDouble,
		RightSeparator:   text.BoxRightSeparatorDouble,
		TopLeft:          text.BoxTopLeftDouble,
		TopRight:         text.BoxTopRightDouble,
		TopSeparator:     text.BoxTopSeparatorDouble,
		UnfinishedRow:    " " + text.BoxUnfinishedLine,
	}

	// StyleBoxLight defines a Boxed-Table like below:
	//  ┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
	//  │   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │   1 │ Arya       │ Stark     │   3000 │                             │
	//  │  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
	//  │ 300 │ Tyrion     │ Lannister │   5000 │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │     │            │ TOTAL     │  10000 │                             │
	//  └─────┴────────────┴───────────┴────────┴─────────────────────────────┘
	StyleBoxLight = BoxStyle{
		BottomLeft:       text.BoxBottomLeft,
		BottomRight:      text.BoxBottomRight,
		BottomSeparator:  text.BoxBottomSeparator,
		Left:             text.BoxLeft,
		LeftSeparator:    text.BoxLeftSeparator,
		MiddleHorizontal: text.BoxHorizontal,
		MiddleSeparator:  text.BoxSeparator,
		MiddleVertical:   text.BoxVertical,
		PaddingLeft:      " ",
		PaddingRight:     " ",
		PageSeparator:    "\n",
		Right:            text.BoxRight,
		RightSeparator:   text.BoxRightSeparator,
		TopLeft:          text.BoxTopLeft,
		TopRight:         text.BoxTopRight,
		TopSeparator:     text.BoxTopSeparator,
		UnfinishedRow:    " " + text.BoxUnfinishedLine,
	}

	// StyleBoxRounded defines a Boxed-Table like below:
	//  ╭─────┬────────────┬───────────┬────────┬─────────────────────────────╮
	//  │   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │   1 │ Arya       │ Stark     │   3000 │                             │
	//  │  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
	//  │ 300 │ Tyrion     │ Lannister │   5000 │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │     │            │ TOTAL     │  10000 │                             │
	//  ╰─────┴────────────┴───────────┴────────┴─────────────────────────────╯
	StyleBoxRounded = BoxStyle{
		BottomLeft:       text.BoxBottomLeftRounded,
		BottomRight:      text.BoxBottomRightRounded,
		BottomSeparator:  text.BoxBottomSeparator,
		Left:             text.BoxLeft,
		LeftSeparator:    text.BoxLeftSeparator,
		MiddleHorizontal: text.BoxHorizontal,
		MiddleSeparator:  text.BoxSeparator,
		MiddleVertical:   text.BoxVertical,
		PaddingLeft:      " ",
		PaddingRight:     " ",
		PageSeparator:    "\n",
		Right:            text.BoxRight,
		RightSeparator:   text.BoxRightSeparator,
		TopLeft:          text.BoxTopLeftRounded,
		TopRight:         text.BoxTopRightRounded,
		TopSeparator:     text.BoxTopSeparator,
		UnfinishedRow:    " " + text.BoxUnfinishedLine,
	}

	// styleBoxTest defines a Boxed-Table like below:
	//  (-----^------------^-----------^--------^-----------------------------)
	//  [<  #>|<FIRST NAME>|<LAST NAME>|<SALARY>|<                           >]
	//  {-----+------------+-----------+--------+-----------------------------}
	//  [<  1>|<Arya      >|<Stark    >|<  3000>|<                           >]
	//  [< 20>|<Jon       >|<Snow     >|<  2000>|<You know nothing, Jon Snow!>]
	//  [<300>|<Tyrion    >|<Lannister>|<  5000>|<                           >]
	//  {-----+------------+-----------+--------+-----------------------------}
	//  [<   >|<          >|<TOTAL    >|< 10000>|<                           >]
	//  \-----v------------v-----------v--------v-----------------------------/
	styleBoxTest = BoxStyle{
		BottomLeft:       "\\",
		BottomRight:      "/",
		BottomSeparator:  "v",
		Left:             "[",
		LeftSeparator:    "{",
		MiddleHorizontal: "-",
		MiddleSeparator:  "+",
		MiddleVertical:   "|",
		PaddingLeft:      "<",
		PaddingRight:     ">",
		PageSeparator:    "\n",
		Right:            "]",
		RightSeparator:   "}",
		TopLeft:          "(",
		TopRight:         ")",
		TopSeparator:     "^",
		UnfinishedRow:    " ~~~",
	}
)

// ColorOptions defines the ANSI colors to use for parts of the Table.
type ColorOptions struct {
	IndexColumn  text.Colors
	Footer       text.Colors
	Header       text.Colors
	Row          text.Colors
	RowAlternate text.Colors
}

var (
	// ColorOptionsDefault defines sensible ANSI color options - basically NONE.
	ColorOptionsDefault = ColorOptions{}

	// ColorOptionsBright renders dark text on bright background.
	ColorOptionsBright = ColorOptionsBlackOnCyanWhite

	// ColorOptionsDark renders bright text on dark background.
	ColorOptionsDark = ColorOptionsCyanWhiteOnBlack

	// ColorOptionsBlackOnBlueWhite renders Black text on Blue/White background.
	ColorOptionsBlackOnBlueWhite = ColorOptions{
		IndexColumn:  text.Colors{text.BgHiBlue, text.FgBlack},
		Footer:       text.Colors{text.BgBlue, text.FgBlack},
		Header:       text.Colors{text.BgHiBlue, text.FgBlack},
		Row:          text.Colors{text.BgHiWhite, text.FgBlack},
		RowAlternate: text.Colors{text.BgWhite, text.FgBlack},
	}

	// ColorOptionsBlackOnCyanWhite renders Black text on Cyan/White background.
	ColorOptionsBlackOnCyanWhite = ColorOptions{
		IndexColumn:  text.Colors{text.BgHiCyan, text.FgBlack},
		Footer:       text.Colors{text.BgCyan, text.FgBlack},
		Header:       text.Colors{text.BgHiCyan, text.FgBlack},
		Row:          text.Colors{text.BgHiWhite, text.FgBlack},
		RowAlternate: text.Colors{text.BgWhite, text.FgBlack},
	}

	// ColorOptionsBlackOnGreenWhite renders Black text on Green/White
	// background.
	ColorOptionsBlackOnGreenWhite = ColorOptions{
		IndexColumn:  text.Colors{text.BgHiGreen, text.FgBlack},
		Footer:       text.Colors{text.BgGreen, text.FgBlack},
		Header:       text.Colors{text.BgHiGreen, text.FgBlack},
		Row:          text.Colors{text.BgHiWhite, text.FgBlack},
		RowAlternate: text.Colors{text.BgWhite, text.FgBlack},
	}

	// ColorOptionsBlackOnMagentaWhite renders Black text on Magenta/White
	// background.
	ColorOptionsBlackOnMagentaWhite = ColorOptions{
		IndexColumn:  text.Colors{text.BgHiMagenta, text.FgBlack},
		Footer:       text.Colors{text.BgMagenta, text.FgBlack},
		Header:       text.Colors{text.BgHiMagenta, text.FgBlack},
		Row:          text.Colors{text.BgHiWhite, text.FgBlack},
		RowAlternate: text.Colors{text.BgWhite, text.FgBlack},
	}

	// ColorOptionsBlackOnRedWhite renders Black text on Red/White background.
	ColorOptionsBlackOnRedWhite = ColorOptions{
		IndexColumn:  text.Colors{text.BgHiRed, text.FgBlack},
		Footer:       text.Colors{text.BgRed, text.FgBlack},
		Header:       text.Colors{text.BgHiRed, text.FgBlack},
		Row:          text.Colors{text.BgHiWhite, text.FgBlack},
		RowAlternate: text.Colors{text.BgWhite, text.FgBlack},
	}

	// ColorOptionsBlackOnYellowWhite renders Black text on Yellow/White
	// background.
	ColorOptionsBlackOnYellowWhite = ColorOptions{
		IndexColumn:  text.Colors{text.BgHiYellow, text.FgBlack},
		Footer:       text.Colors{text.BgYellow, text.FgBlack},
		Header:       text.Colors{text.BgHiYellow, text.FgBlack},
		Row:          text.Colors{text.BgHiWhite, text.FgBlack},
		RowAlternate: text.Colors{text.BgWhite, text.FgBlack},
	}

	// ColorOptionsBlueWhiteOnBlack renders Blue/White text on Black background.
	ColorOptionsBlueWhiteOnBlack = ColorOptions{
		IndexColumn:  text.Colors{text.FgHiBlue, text.BgHiBlack},
		Footer:       text.Colors{text.FgBlue, text.BgHiBlack},
		Header:       text.Colors{text.FgHiBlue, text.BgHiBlack},
		Row:          text.Colors{text.FgHiWhite, text.BgBlack},
		RowAlternate: text.Colors{text.FgWhite, text.BgBlack},
	}

	// ColorOptionsCyanWhiteOnBlack renders Cyan/White text on Black background.
	ColorOptionsCyanWhiteOnBlack = ColorOptions{
		IndexColumn:  text.Colors{text.FgHiCyan, text.BgHiBlack},
		Footer:       text.Colors{text.FgCyan, text.BgHiBlack},
		Header:       text.Colors{text.FgHiCyan, text.BgHiBlack},
		Row:          text.Colors{text.FgHiWhite, text.BgBlack},
		RowAlternate: text.Colors{text.FgWhite, text.BgBlack},
	}

	// ColorOptionsGreenWhiteOnBlack renders Green/White text on Black
	// background.
	ColorOptionsGreenWhiteOnBlack = ColorOptions{
		IndexColumn:  text.Colors{text.FgHiGreen, text.BgHiBlack},
		Footer:       text.Colors{text.FgGreen, text.BgHiBlack},
		Header:       text.Colors{text.FgHiGreen, text.BgHiBlack},
		Row:          text.Colors{text.FgHiWhite, text.BgBlack},
		RowAlternate: text.Colors{text.FgWhite, text.BgBlack},
	}

	// ColorOptionsMagentaWhiteOnBlack renders Magenta/White text on Black
	// background.
	ColorOptionsMagentaWhiteOnBlack = ColorOptions{
		IndexColumn:  text.Colors{text.FgHiMagenta, text.BgHiBlack},
		Footer:       text.Colors{text.FgMagenta, text.BgHiBlack},
		Header:       text.Colors{text.FgHiMagenta, text.BgHiBlack},
		Row:          text.Colors{text.FgHiWhite, text.BgBlack},
		RowAlternate: text.Colors{text.FgWhite, text.BgBlack},
	}

	// ColorOptionsRedWhiteOnBlack renders Red/White text on Black background.
	ColorOptionsRedWhiteOnBlack = ColorOptions{
		IndexColumn:  text.Colors{text.FgHiRed, text.BgHiBlack},
		Footer:       text.Colors{text.FgRed, text.BgHiBlack},
		Header:       text.Colors{text.FgHiRed, text.BgHiBlack},
		Row:          text.Colors{text.FgHiWhite, text.BgBlack},
		RowAlternate: text.Colors{text.FgWhite, text.BgBlack},
	}

	// ColorOptionsYellowWhiteOnBlack renders Yellow/White text on Black
	// background.
	ColorOptionsYellowWhiteOnBlack = ColorOptions{
		IndexColumn:  text.Colors{text.FgHiYellow, text.BgHiBlack},
		Footer:       text.Colors{text.FgYellow, text.BgHiBlack},
		Header:       text.Colors{text.FgHiYellow, text.BgHiBlack},
		Row:          text.Colors{text.FgHiWhite, text.BgBlack},
		RowAlternate: text.Colors{text.FgWhite, text.BgBlack},
	}
)

// FormatOptions defines the text-formatting to perform on parts of the Table.
type FormatOptions struct {
	Footer text.Format
	Header text.Format
	Row    text.Format
}

var (
	// FormatOptionsDefault defines sensible formatting options.
	FormatOptionsDefault = FormatOptions{
		Footer: text.FormatUpper,
		Header: text.FormatUpper,
		Row:    text.FormatDefault,
	}
)

// Options defines the global options that determine how the Table is
// rendered.
type Options struct {
	// DrawBorder enables or disables drawing the border around the Table.
	// Example of a table where it is disabled:
	//     # │ FIRST NAME │ LAST NAME │ SALARY │
	//  ─────┼────────────┼───────────┼────────┼─────────────────────────────
	//     1 │ Arya       │ Stark     │   3000 │
	//    20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow!
	//   300 │ Tyrion     │ Lannister │   5000 │
	//  ─────┼────────────┼───────────┼────────┼─────────────────────────────
	//       │            │ TOTAL     │  10000 │
	DrawBorder bool

	// SeparateColumns enables or disable drawing border between columns.
	// Example of a table where it is disabled:
	//  ┌─────────────────────────────────────────────────────────────────┐
	//  │   #  FIRST NAME  LAST NAME  SALARY                              │
	//  ├─────────────────────────────────────────────────────────────────┤
	//  │   1  Arya        Stark        3000                              │
	//  │  20  Jon         Snow         2000  You know nothing, Jon Snow! │
	//  │ 300  Tyrion      Lannister    5000                              │
	//  │                  TOTAL       10000                              │
	//  └─────────────────────────────────────────────────────────────────┘
	SeparateColumns bool

	// SeparateFooter enables or disable drawing border between the footer and
	// the rows. Example of a table where it is disabled:
	//  ┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
	//  │   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │   1 │ Arya       │ Stark     │   3000 │                             │
	//  │  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
	//  │ 300 │ Tyrion     │ Lannister │   5000 │                             │
	//  │     │            │ TOTAL     │  10000 │                             │
	//  └─────┴────────────┴───────────┴────────┴─────────────────────────────┘
	SeparateFooter bool

	// SeparateHeader enables or disable drawing border between the header and
	// the rows. Example of a table where it is disabled:
	//  ┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
	//  │   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
	//  │   1 │ Arya       │ Stark     │   3000 │                             │
	//  │  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
	//  │ 300 │ Tyrion     │ Lannister │   5000 │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │     │            │ TOTAL     │  10000 │                             │
	//  └─────┴────────────┴───────────┴────────┴─────────────────────────────┘
	SeparateHeader bool

	// SeparateRows enables or disables drawing separators between each row.
	// Example of a table where it is enabled:
	//  ┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
	//  │   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │   1 │ Arya       │ Stark     │   3000 │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │ 300 │ Tyrion     │ Lannister │   5000 │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │     │            │ TOTAL     │  10000 │                             │
	//  └─────┴────────────┴───────────┴────────┴─────────────────────────────┘
	SeparateRows bool
}

var (
	// OptionsDefault defines sensible global options.
	OptionsDefault = Options{
		DrawBorder:      true,
		SeparateColumns: true,
		SeparateFooter:  true,
		SeparateHeader:  true,
		SeparateRows:    false,
	}

	// OptionsNoBorders sets up a table without any borders.
	OptionsNoBorders = Options{
		DrawBorder:      false,
		SeparateColumns: true,
		SeparateFooter:  true,
		SeparateHeader:  true,
		SeparateRows:    false,
	}

	// OptionsNoBordersAndSeparators sets up a table without any borders or
	// separators.
	OptionsNoBordersAndSeparators = Options{
		DrawBorder:      false,
		SeparateColumns: false,
		SeparateFooter:  false,
		SeparateHeader:  false,
		SeparateRows:    false,
	}
)
