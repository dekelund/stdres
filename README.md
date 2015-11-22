# StdRes

Print action results to STDOUT

For license information, see LICENSE.

## Introduction
Package stdres provides the means to select colors for STDOUT based on action results.
By using this package you are able to print text to the buffer, and postpone the color
selection until action has finished.

## What can StdRes be used for?
If you write a compiler or test framework using colorized output, but don't want to deal
with text later, instead you want to record text before action has finished and reduce
if-statements and clutter in your code.

Note that this library will not update already printed text, it only buffers output till
later, when the result is known.

## Installation

`go get github.com/dekelund/stdres`

## Usage

See documentation and test file for more information.
