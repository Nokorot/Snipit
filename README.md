
This is a snippets script, intended to be used in vim. It may be used to 
insert templates/snippets by executing this script inside vim with the following command:
    :r!snipits tex<CR>?<CSR><CR>vf>
Note: here 'tex' indicates that the snippets in the folder 'tex' are to be used.

The script uses demnu to select a snippet, by a search.

DEPENDENCIES:

 - My version of [dmenu](https://github.com/Nokorot/dmenu), as it relies on 
    the json patch for dmenu as well as a flag [-jd jsondepth], which limits the depth of the search and instead returns the selected object as a new json text.
