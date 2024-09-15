module view

go 1.23

require (
    model v0.0.0
    service v0.0.0
)

replace model v0.0.0 => "../model"
replace service v0.0.0 => "../service"